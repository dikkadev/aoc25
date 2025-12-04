#!/bin/bash

# Exit on any error
set -euo pipefail

LOG_FILE="create_day.log"
TEMPLATE_FILE="days/template.go"
BUILD_EXCLUDE="//go:build ignore"

log() {
    # echo "$(date +'%Y-%m-%d %H:%M:%S') - $1" | tee -a "$LOG_FILE"
    echo "$(date +'%Y-%m-%d %H:%M:%S') - $1"
}

add_import_to_main() {
    local day="$1"
    local import_line='_ "github.com/dikkadev/aoc25/days/'$day'"'
    local main_file="main.go"
    
    # Check if main.go exists
    if [[ ! -f "$main_file" ]]; then
        log "ERROR: main.go file not found"
        return 1
    fi
    
    # Check if import already exists
    if grep -q "$import_line" "$main_file"; then
        log "Import already exists in main.go: $import_line"
        return 0
    fi
    
    # Add import after the existing days import
    if sed -i '/github.com\/dikkadev\/aoc25\/days/a\
\t'"$import_line" "$main_file"; then
        log "Added import to main.go: $import_line"
        return 0
    else
        log "ERROR: Failed to add import to main.go"
        return 1
    fi
}

usage() {
    echo "Usage: $0 [-gold] <day>"
    exit 1
}

if [[ $# -lt 1 || $# -gt 2 ]]; then
    usage
fi

GOLD_MODE=false
if [[ "$1" == "-gold" ]]; then
    GOLD_MODE=true
    shift
fi

INPUT="$1"
DAY_DIR="days/$INPUT"
INPUT_FILE="$DAY_DIR/$INPUT.go"
GOLD_FILE="$DAY_DIR/${INPUT}_gold.go"
INPUT_DIR="input"
INPUT_INPUT_FILE="$INPUT_DIR/$(printf "%02d" "$INPUT").input"
INPUT_SMALL_FILE="$INPUT_DIR/$(printf "%02d" "$INPUT")_small.input"

if $GOLD_MODE; then
    # Ensure the base script has already been run for the given input
    if [[ ! -f "$INPUT_FILE" ]]; then
        log "ERROR: Base setup for day '$INPUT' is not complete. Run the script without -gold first."
        exit 1
    fi

    # Copy the INPUT.go to INPUT_gold.go
    if [[ ! -f "$GOLD_FILE" ]]; then
        cp "$INPUT_FILE" "$GOLD_FILE"
        log "Created gold file: $GOLD_FILE"
    else
        log "Gold file already exists: $GOLD_FILE"
    fi

    # Add build exclude line to the top of INPUT.go if not already present
    if ! grep -q "$BUILD_EXCLUDE" "$INPUT_FILE"; then
        sed -i "1s|^|$BUILD_EXCLUDE\n\n|" "$INPUT_FILE"
        log "Added build exclude line to: $INPUT_FILE"
    else
        log "Build exclude line already present in: $INPUT_FILE"
    fi

    log "Gold mode completed successfully for day: $INPUT"
    exit 0
fi

log "Script started for day: $INPUT"

# Create day directory
if [[ ! -d "$DAY_DIR" ]]; then
    mkdir -p "$DAY_DIR"
    log "Created directory: $DAY_DIR"
else
    log "Directory already exists: $DAY_DIR"
fi

# Create INPUT.go file from template
if [[ -f "$TEMPLATE_FILE" ]]; then
    if [[ ! -f "$INPUT_FILE" ]]; then
        tail -n +2 "$TEMPLATE_FILE" > "$INPUT_FILE"
        log "Created file: $INPUT_FILE from template (without first line)"
        
        # Update the DAY constant to the actual day number
        sed -i "s/const DAY = 0/const DAY = $INPUT/" "$INPUT_FILE"
        log "Updated DAY constant to: $INPUT"
        
        # Add import to main.go
        if ! add_import_to_main "$INPUT"; then
            log "ERROR: Failed to add import to main.go"
            exit 1
        fi
    else
        log "File already exists: $INPUT_FILE"
    fi
else
    log "ERROR: Template file not found: $TEMPLATE_FILE"
    exit 1
fi

# Create input directory
if [[ ! -d "$INPUT_DIR" ]]; then
    mkdir -p "$INPUT_DIR"
    log "Created directory: $INPUT_DIR"
else
    log "Directory already exists: $INPUT_DIR"
fi

# Create empty input files
for FILE in "$INPUT_INPUT_FILE" "$INPUT_SMALL_FILE"; do
    if [[ ! -f "$FILE" ]]; then
        touch "$FILE"
        log "Created empty file: $FILE"
    else
        log "File already exists: $FILE"
    fi
done

log "Script completed successfully for day: $INPUT"
