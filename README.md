# AoC 2025 Framework

A Go framework for solving Advent of Code 2025 puzzles.

## Structure

```
aoc25/
├── main.go          # Main entry point
├── days/            # Day implementations
│   ├── days.go      # Day registration and management
│   └── template.go  # Template for new days
├── input/           # Input handling
│   └── input.go     # Input parsing utilities
├── new.sh           # Script to create new days
└── README.md        # This file
```

## Usage

### Creating a New Day

Use the provided script to create a new day:

```bash
./new.sh 1
```

This creates:
- `days/1/1.go` - Day implementation file
- `input/1.input` - Real puzzle input
- `input/1_small.input` - Test/small input

### Running a Day

Run a specific day with real input:
```bash
go run . -d 1
```

Run with small/test input:
```bash
go run . -d 1 -s
```

### Implementing a Day

1. Edit the created file `days/1/1.go`
2. Set the `DAY` constant to the day number
3. Implement the `Solve` function
4. The function receives:
   - `*input.Input` - Parsed input data
   - `*slog.Logger` - Logger for debugging
5. Return the puzzle answer as an integer

### Input Handling

The `input.Input` struct provides several methods:
- `AugmentedLineStream()` - Stream lines with metadata
- `Lines()` - Get all lines as a slice
- `String()` - Get entire input as a string

### Gold Mode

When you want to lock in a solution and create a gold version:

```bash
./new.sh -gold 1
```

This:
- Creates `days/1/1_gold.go` (copy of your solution)
- Adds `//go:build ignore` to `days/1/1.go`

## Features

- **Automatic day registration** - Days are registered via `init()` functions
- **Structured logging** - Each day gets its own logger
- **Input management** - Handles both small and real inputs
- **CPU profiling** - Built-in profiling for performance analysis
- **Clean separation** - Each day in its own directory

## Dependencies

- Go 1.23.0+
- `github.com/dikkadev/prettyslog` for pretty logging

## Example Day Implementation

```go
package day

import (
    "log/slog"
    "github.com/dikkadev/aoc25/days"
    "github.com/dikkadev/aoc25/input"
)

const DAY = 1

func init() {
    days.RegisterDay(DAY, Solve)
}

func Solve(input *input.Input, log *slog.Logger) int {
    result := 0
    for l := range input.AugmentedLineStream() {
        if len(l.T) == 0 {
            continue
        }
        // Your solution logic here
        log.Debug(l.T)
    }
    return result
}
```

## Tips

- Use the `-s` flag for debugging with small inputs
- Each day gets its own logger prefixed with "D##"
- The framework automatically handles input file paths
- Use `input.AugmentedLineStream()` for line-by-line processing with line numbers