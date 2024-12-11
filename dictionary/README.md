# Dictionary Word Processor

This script processes a list of English words, fetches their definitions using an online API, and saves the results into CSV files organized by their starting letters.

---

## Features

- Fetches English word definitions from the [Free Dictionary API](https://dictionaryapi.dev/).
- Saves results in CSV format, organized by starting letter.
- Configurable using a `.env` file or command-line arguments.
- Supports parallel processing for faster execution.
- Allows filtering words by starting letters and excluding specific letters.
- Disk space checks to avoid running out of storage during processing.
- Optional test mode to debug without saving files.

---

## Prerequisites

1. **Python 3.x**  
   Ensure Python 3 is installed. Verify by running:
   ```bash
   python3 --version
   ```

2. **Install Dependencies**  
   Install required Python packages:
   ```bash
   pip install requests python-dotenv
   ```

---

## Setup

### **1. Clone or Download the Repository**
Save the script to your local machine.

### **2. Create a `.env` File**
Create a file named `.env` in the script directory and add the following configuration:

```ini
LIST_WORDS_URL=https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt
API_URL=https://api.dictionaryapi.dev/api/v2/entries/en/
OUTPUT_DIR=dictionary_words_by_letter
LOG_FILE=process_log.txt
DISK_SPACE_THRESHOLD_GB=5
REQUESTS_TIMEOUT=10
RETRIES=5
ALWAYS_CONFIRM=true
```

Modify values as necessary.

---

## Usage

Run the script using `python3`:

```bash
python3 dump_dictionary.py [OPTIONS]
```

### **Available Options**
| Argument            | Description                                                                        | Default                      |
|---------------------|------------------------------------------------------------------------------------|------------------------------|
| `--only-startwith`  | Process words starting with specific letters (regex pattern, e.g., `a\|b\|[c-d]`). | `.*` (all letters)           |
| `--exclude-letters` | Exclude words starting with specific letters (e.g., `x\|y\|z`).                    | None                         |
| `--start-after`     | Start processing after a specific letter (e.g., `m`).                              | None                         |
| `--output-dir`      | Directory to save CSV files.                                                       | `dictionary_words_by_letter` |
| `--log-file`        | Log file location.                                                                 | `process_log.txt`            |
| `--timeout`         | Request timeout in seconds.                                                        | `10`                         |
| `--retries`         | Number of retries for fetching definitions.                                        | `5`                          |
| `--confirm`         | Ask for confirmation before processing.                                            | As set in `.env`             |
| `--skip-disk-check` | Skip disk space checks.                                                            | Disabled                     |
| `--limit-words`     | Limit the total number of words to process.                                        | None                         |
| `--overwrite`       | Overwrite existing CSV files.                                                      | Disabled                     |
| `--test-mode`       | Run the script without saving files (for debugging).                               | Disabled                     |

---

## Examples

1. **Process all words**:
   ```bash
   python3 dump_dictionary.py
   ```

2. **Process words starting with specific letters**:
   ```bash
   python3 dump_dictionary.py --only-startwith "a|b|c"
   ```

3. **Exclude certain letters**:
   ```bash
   python3 dump_dictionary.py --exclude-letters "x|y|z"
   ```

4. **Limit the number of words**:
   ```bash
   python3 dump_dictionary.py --limit-words 1000
   ```

5. **Run in test mode**:
   ```bash
   python3 dump_dictionary.py --test-mode
   ```


## Output

1. **Processed CSV Files**  
   Saved in the directory specified by the `--output-dir` argument or the `OUTPUT_DIR` in the `.env` file.

2. **Log File**  
   Logs are written to the file specified by the `--log-file` argument or the `LOG_FILE` in the `.env` file.

---

