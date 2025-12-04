#!/bin/bash

# Exit on any error
set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

log() {
    echo -e "${BLUE}[$(date +'%Y-%m-%d %H:%M:%S')]${NC} $1"
}

error() {
    echo -e "${RED}[ERROR]${NC} $1" >&2
}

success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

usage() {
    echo "Usage: $0 [-gold] <day>"
    echo "  -gold    Create gold version (locks in solution)"
    echo "  <day>    Day number (1-25)"
    exit 1
}

# Check if we're in the right directory
if [[ ! -f "app/build.gradle.kts" ]]; then
    error "Must be run from the aoc25 project root"
    exit 1
fi

# Parse arguments
GOLD_MODE=false
DAY=""

if [[ $# -lt 1 || $# -gt 2 ]]; then
    usage
fi

if [[ "$1" == "-gold" ]]; then
    GOLD_MODE=true
    shift
fi

DAY="$1"

# Validate day number
if ! [[ "$DAY" =~ ^[1-9]$|^1[0-9]$|^2[0-5]$ ]]; then
    error "Day must be between 1 and 25"
    exit 1
fi

# Set up paths
DAY_DIR="app/src/main/kotlin/com/aoc25/days"
DAY_FILE="$DAY_DIR/Day${DAY}.kt"
GOLD_FILE="$DAY_DIR/Day${DAY}_Gold.kt"
INPUT_DIR="app/src/main/resources/input"
INPUT_FILE="$INPUT_DIR/${DAY}.input"
SMALL_INPUT_FILE="$INPUT_DIR/${DAY}_small.input"

log "Creating AoC Day $DAY (Gold mode: $GOLD_MODE)"

# Gold mode logic
if $GOLD_MODE; then
    # Check if the base day file exists
    if [[ ! -f "$DAY_FILE" ]]; then
        error "Base day file not found: $DAY_FILE"
        error "Run without -gold first to create the day"
        exit 1
    fi
    
    # Create gold version if it doesn't exist
    if [[ ! -f "$GOLD_FILE" ]]; then
        cp "$DAY_FILE" "$GOLD_FILE"
        # Update the class name in gold file
        sed -i "s/class Day${DAY}/class Day${DAY}_Gold/" "$GOLD_FILE"
        success "Created gold file: $GOLD_FILE"
    else
        warning "Gold file already exists: $GOLD_FILE"
    fi
    
    # Comment out the original day class
    if ! grep -q "// GOLD MODE: Day ${DAY} is locked" "$DAY_FILE"; then
        sed -i "1i// GOLD MODE: Day ${DAY} is locked - see Day${DAY}_Gold.kt" "$DAY_FILE"
        sed -i "s/@AoCDay(number = $DAY)/\/\/ @AoCDay(number = $DAY)/" "$DAY_FILE"
        sed -i "s/^class Day${DAY}/\/\/ class Day${DAY}/" "$DAY_FILE"
        sed -i "s/^    override val dayNumber/\/\/     override val dayNumber/" "$DAY_FILE"
        sed -i "s/^    override suspend fun solve/\/\/     override suspend fun solve/" "$DAY_FILE"
        success "Locked base day file: $DAY_FILE"
    else
        warning "Base day file already locked: $DAY_FILE"
    fi
    
    success "Gold mode completed for day $DAY"
    exit 0
fi

# Regular day creation
# Create day file from template if it doesn't exist
if [[ ! -f "$DAY_FILE" ]]; then
    # Copy template and customize
    sed "s/DayTemplate/Day${DAY}/g; s/number = 0/number = $DAY/g; s/dayNumber: Int = 0/dayNumber: Int = $DAY/g" \
        "$DAY_DIR/DayTemplate.kt" > "$DAY_FILE"
    
    success "Created day file: $DAY_FILE"
else
    warning "Day file already exists: $DAY_FILE"
fi

# Create input directory if it doesn't exist
if [[ ! -d "$INPUT_DIR" ]]; then
    mkdir -p "$INPUT_DIR"
    success "Created input directory: $INPUT_DIR"
fi

# Create empty input files
for file in "$INPUT_FILE" "$SMALL_INPUT_FILE"; do
    if [[ ! -f "$file" ]]; then
        touch "$file"
        success "Created input file: $file"
    else
        warning "Input file already exists: $file"
    fi
done

# Update the DayRegistry to include the new day
REGISTRY_FILE="app/src/main/kotlin/com/aoc25/framework/DayRegistry.kt"
if [[ -f "$REGISTRY_FILE" ]]; then
    # Check if the day is already in the known list
    if ! grep -q "com.aoc25.days.Day${DAY}" "$REGISTRY_FILE"; then
        # Add the new day to the known list
        sed -i "/knownDayClasses = listOf(/a\\                \"com.aoc25.days.Day${DAY}\"," "$REGISTRY_FILE"
        success "Added Day${DAY} to registry discovery list"
    else
        warning "Day${DAY} already in registry discovery list"
    fi
    
    # For gold mode, also add the gold class
    if $GOLD_MODE; then
        if ! grep -q "com.aoc25.days.Day${DAY}_Gold" "$REGISTRY_FILE"; then
            sed -i "/\"com.aoc25.days.Day${DAY}\",/a\\                \"com.aoc25.days.Day${DAY}_Gold\"," "$REGISTRY_FILE"
            success "Added Day${DAY}_Gold to registry discovery list"
        else
            warning "Day${DAY}_Gold already in registry discovery list"
        fi
    fi
fi

success "Day $DAY setup completed!"
echo
echo "Next steps:"
echo "1. Edit $DAY_FILE to implement your solution"
echo "2. Add test input to $SMALL_INPUT_FILE"
echo "3. Add real input to $INPUT_FILE"
echo "4. Run with: ./gradlew run --args=\"-d $DAY -s\" (for small input)"
echo "5. Run with: ./gradlew run --args=\"-d $DAY\" (for real input)"