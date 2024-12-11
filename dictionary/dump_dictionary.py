import requests
import csv
import os
import shutil
import re
import argparse
import time
from dotenv import load_dotenv
from string import ascii_lowercase, ascii_uppercase

load_dotenv()

LIST_WORDS_URL = os.getenv("LIST_WORDS_URL")
API_URL = os.getenv("API_URL")
OUTPUT_DIR = os.getenv("OUTPUT_DIR", "output")
LOG_FILE = os.getenv("LOG_FILE", "dictionary.log")
DISK_SPACE_THRESHOLD_GB = int(os.getenv("DISK_SPACE_THRESHOLD_GB", 1))
REQUESTS_TIMEOUT = int(os.getenv("REQUESTS_TIMEOUT", 10))
RETRIES = int(os.getenv("RETRIES", 3))
ALWAYS_CONFIRM = os.getenv("ALWAYS_CONFIRM", "false").lower() == "true"

response = requests.get(LIST_WORDS_URL)
if response.status_code != 200:
    print(f"Failed to fetch words list. Status code: {response.status_code}")
    exit(1)
words = response.text.split("\n")

def fetch_word_definition_with_retries(word, retries=RETRIES, timeout=REQUESTS_TIMEOUT):
    delay = 1
    for attempt in range(retries):
        try:
            _response = requests.get(API_URL + word, timeout=timeout)
            if _response.status_code == 200:
                data = _response.json()
                return data[0]["meanings"][0]["definitions"][0]["definition"]
            elif _response.status_code == 429:
                time.sleep(delay)
                delay *= 2
            else:
                return f"Error: {_response.status_code}"
        except Exception as e:
            return f"Error: {str(e)}"
    return "Failed after retries."

def check_disk_space(threshold_gb=DISK_SPACE_THRESHOLD_GB):
    total, used, free = shutil.disk_usage("/")
    free_gb = free // (2**30)
    return free_gb, free_gb >= threshold_gb

def log_message(message, letter=None):
    print(message)
    with open(LOG_FILE, "a") as log:
        if letter:
            log.write(f"[{letter}] {message}\n")
        else:
            log.write(message + "\n")

parser = argparse.ArgumentParser(description="Fetch word definitions and save to CSV.")
parser.add_argument("--only-startwith", type=str, default=".*", help="Regex pattern to filter starting letters (e.g., 'ac|b|[e-g]').")
parser.add_argument("--output-dir", type=str, default=OUTPUT_DIR, help="Directory to save CSV files.")
parser.add_argument("--log-file", type=str, default=LOG_FILE, help="Log file location.")
parser.add_argument("--timeout", type=int, default=REQUESTS_TIMEOUT, help="Request timeout in seconds.")
parser.add_argument("--retries", type=int, default=RETRIES, help="Number of retries for fetching definitions.")
parser.add_argument("--confirm", action="store_true", default=ALWAYS_CONFIRM, help="Always confirm before processing.")
parser.add_argument("--skip-disk-check", action="store_true", help="Skip disk space check.")
parser.add_argument("--limit-words", type=int, help="Limit the total number of words to process per letter.")
parser.add_argument("--max-words", type=int, help="Global limit on the total number of words to process.")
parser.add_argument("--overwrite", action="store_true", help="Overwrite existing CSV files.")
parser.add_argument("--exclude-letters", type=str, help="Exclude letters from processing (e.g., 'x|y|z').")
parser.add_argument("--start-after", type=str, help="Start processing after a specific letter.")
parser.add_argument("--test-mode", action="store_true", help="Run the script in test mode without saving files.")
args = parser.parse_args()

pattern = re.compile(f"^{args.only_startwith}")
exclude_letters = set(args.exclude_letters.split("|")) if args.exclude_letters else set()
start_after = args.start_after.lower() if args.start_after else None

dictionary_words_by_letter = {
    letter: [word for word in words if pattern.match(word) and word.startswith(letter)]
    for letter in ascii_lowercase + ascii_uppercase
    if letter.lower() not in exclude_letters
}

if start_after:
    dictionary_words_by_letter = {
        letter: words_list
        for letter, words_list in dictionary_words_by_letter.items()
        if letter.lower() > start_after
    }

total_words = sum(len(words_list) for words_list in dictionary_words_by_letter.values())
available_space, has_enough_space = check_disk_space() if not args.skip_disk_check else (None, True)

print(f"Total words to process: {total_words}")
if not args.skip_disk_check:
    print(f"Available disk space: {available_space}GB")

if not has_enough_space:
    print("Insufficient disk space. Process cannot continue.")
    exit(1)

if not ALWAYS_CONFIRM and args.confirm:
    confirmation = input("Do you want to proceed? (yes/no): ").strip().lower()
    if confirmation != "yes":
        print("Process aborted by user.")
        exit(0)

os.makedirs(args.output_dir, exist_ok=True)

total_words_processed = 0

def process_letter(letter, words_list):
    global total_words_processed
    if not words_list:
        return
    log_message(f"Processing words starting with letter '{letter}'", letter)
    file_path = os.path.join(args.output_dir, f"{letter}_words.csv")
    if os.path.exists(file_path) and not args.overwrite:
        log_message(f"Skipping existing file '{file_path}'", letter)
        return
    with open(file_path, mode="w", newline="") as file:
        writer = csv.writer(file)
        writer.writerow(["Word", "Definition"])
        for word in words_list:
            if args.limit_words and len(words_list) > args.limit_words:
                words_list = words_list[:args.limit_words]
            if args.max_words and total_words_processed >= args.max_words:
                return
            definition = fetch_word_definition_with_retries(word, retries=args.retries, timeout=args.timeout)
            if not args.test_mode:
                writer.writerow([word, definition])
            log_message(f"Saved word '{word}' with definition.", letter)
            total_words_processed += 1

for letter, words_list in dictionary_words_by_letter.items():
    process_letter(letter, words_list)
    if args.max_words and total_words_processed >= args.max_words:
        break

log_message("Processing completed.")
