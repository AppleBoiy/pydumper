#!/bin/bash

search_words=("example" "web scraping" "python tutorial")
num_results=5
max_threads=2
get_running_threads() {
    echo $(ps aux | grep -c 'python google_scrap.py')
}
for word in "${search_words[@]}"; do
    while [ $(get_running_threads) -ge $max_threads ]; do
        sleep 1
    done

    python google_scrap.py "$word" --num_results "$num_results" &
done

wait
echo "All searches and scrapes are complete!"