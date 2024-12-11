import argparse
import os
import time
import urllib.parse
import re
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.common.exceptions import TimeoutException

base_dir = "web_scraping"
chromedriver_path = '/usr/local/bin/chromedriver'

def to_snake_case(s):
    return re.sub(r'[^a-zA-Z0-9]', '_', s).lower()

def fetch_google_search_results(query, num_results=10):
    search_results = []
    url = f"https://www.google.com/search?q={urllib.parse.quote(query)}"

    chrome_options = webdriver.ChromeOptions()
    chrome_options.add_argument("--headless")
    chrome_options.add_argument("--no-sandbox")
    chrome_options.add_argument("--disable-dev-shm-usage")
    chrome_options.add_argument("--remote-debugging-port=9222")
    chrome_options.add_argument("--disable-gpu")

    driver = webdriver.Chrome(options=chrome_options)

    try:
        driver.get(url)
        WebDriverWait(driver, 10).until(EC.presence_of_element_located((By.CSS_SELECTOR, 'h3')))
        links = driver.find_elements(By.CSS_SELECTOR, 'a[jsname="UWckNb"]')

        for link in links:
            href = link.get_attribute('href')
            if href and href.startswith('http'):
                search_results.append(href)

            if len(search_results) >= num_results:
                break

    except TimeoutException:
        print(f"Timeout while fetching search results for '{query}'. Skipping...")

    finally:
        driver.quit()

    return search_results

def save_to_file(content, filename):
    with open(filename, 'w', encoding='utf-8') as file:
        file.write(content)

def main(search_words, num_results):
    for word in search_words:
        print(f"Fetching search results for '{word}'...")
        urls = fetch_google_search_results(word, num_results)

        os.makedirs(f"{base_dir}/{word[0]}/{to_snake_case(word)}", exist_ok=True)

        with open(f"{base_dir}/{word[0]}/{to_snake_case(word)}/label.csv", "w") as f:
            f.write("word,url,file\n")

        label_file = f"{base_dir}/{word[0]}/{to_snake_case(word)}/label.csv"
        label = open(label_file, "a")

        for idx, url in enumerate(urls):
            print(f"Scraping page {idx + 1}: {url}")

            chrome_options = webdriver.ChromeOptions()
            chrome_options.add_argument("--headless")
            chrome_options.add_argument("--no-sandbox")
            chrome_options.add_argument("--disable-dev-shm-usage")
            chrome_options.add_argument("--remote-debugging-port=9222")
            chrome_options.add_argument("--disable-gpu")

            driver = webdriver.Chrome(options=chrome_options)
            os.makedirs(f"{base_dir}/{word[0]}/{to_snake_case(word)}", exist_ok=True)

            try:
                driver.set_page_load_timeout(10)
                driver.get(url)
                WebDriverWait(driver, 10).until(EC.presence_of_element_located((By.TAG_NAME, "body")))
                page_content = driver.page_source

                filename = f"{base_dir}/{word[0]}/{to_snake_case(word)}/{idx + 1}.html"
                label.write(f"{word},{url},{idx + 1}\n")

                save_to_file(page_content, filename)
                print(f"Saved content from {url} to {filename}")

            except TimeoutException:
                print(f"Timeout occurred while loading the page: {url}. Skipping...")

            finally:
                driver.quit()

        label.close()

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Scrape Google search results for given search terms.")
    parser.add_argument("search_words", metavar="word", type=str, nargs="+", help="List of words to search for")
    parser.add_argument("--num_results", type=int, default=10, help="Number of results per word to fetch")

    args = parser.parse_args()

    main(args.search_words, args.num_results)