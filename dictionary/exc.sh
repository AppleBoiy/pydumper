#!/bin/bash

LOG_FILE="dictionary_full_run.log"
OUTPUT_DIR="output"
SCRIPT="dump_dictionary.py"
LETTERS=$(echo "abc" | sed 's/./& /g')

: > $LOG_FILE

for letter in $LETTERS; do
    echo "Processing words starting with: $letter" | tee -a $LOG_FILE
    nohup python3 $SCRIPT \
        --only-startwith "^$letter" \
        --output-dir $OUTPUT_DIR \
        --log-file $LOG_FILE \
        --overwrite \
        --skip-disk-check \
        --max-words 1 \
        --confirm >> $LOG_FILE 2>&1 &
done

echo "All letters processed successfully. Logs saved in $LOG_FILE."