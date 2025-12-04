# AoC 2025 - Kotlin Framework

A modern Kotlin framework for solving Advent of Code 2025 puzzles, built with Java 25 and the latest Kotlin features.

## Features

- 🎯 **Java 25** - Latest Java features and performance
- ⚡ **Kotlin Coroutines** - Efficient streaming for large inputs
- 🔄 **Automatic Day Discovery** - Reflection-based day registration
- 📝 **Structured Logging** - SLF4J with Logback
- 🛠️ **CLI Interface** - Built with Clikt for easy command-line usage
- 📦 **Type Safety** - Sealed classes and data classes
- 🎨 **Idiomatic Kotlin** - Extension functions, flows, and more

## Quick Start

### Prerequisites

- Java 25
- Gradle (included via wrapper)

### Creating a New Day

```bash
./scripts/new-day.sh 1
```

This creates:
- `app/src/main/kotlin/com/aoc25/days/Day1.kt` - Day implementation
- `app/src/main/resources/input/1.input` - Real puzzle input
- `app/src/main/resources/input/1_small.input` - Test/small input

### Running a Day

With small/test input:
```bash
./gradlew run --args="-d 1 -s"
```

With real input:
```bash
./gradlew run --args="-d 1"
```

With debug logging:
```bash
./gradlew run --args="-d 1 -s --debug"
```

### Implementing a Day

Edit the generated `Day1.kt` file:

```kotlin
@AoCDay(number = 1)
class Day1 : Day {
    
    override val dayNumber: Int = 1
    
    override suspend fun solve(input: Input, logger: Logger): Any {
        var result = 0
        
        // Process lines with coroutines
        input.augmentedLineFlow().collect { line ->
            if (line.text.isNotEmpty()) {
                logger.debug("Line ${line.lineNumber}: ${line.text}")
                // Your solution logic here
            }
        }
        
        return result
    }
}
```

## Input Processing

The `Input` class provides multiple ways to process puzzle data:

```kotlin
// Stream lines (memory efficient)
input.lineFlow().collect { line -> /* process */ }

// Stream lines with line numbers
input.augmentedLineFlow().collect { augmentedLine ->
    val line = augmentedLine.text
    val lineNumber = augmentedLine.lineNumber
}

// Get all lines at once
val lines = input.lines()

// Stream words
input.wordFlow().collect { word -> /* process */ }

// Stream characters
input.charFlow().collect { char -> /* process */ }

// Get entire input as string
val text = input.asString()
```

## Gold Mode

Lock in a solution by creating a gold version:

```bash
./scripts/new-day.sh -gold 1
```

This:
- Creates `Day1_Gold.kt` (copy of your solution)
- Disables the original `Day1.kt` class
- Allows you to experiment while preserving the working solution

## Project Structure

```
aoc25/
├── app/build.gradle.kts          # Gradle build configuration
├── app/src/main/kotlin/com/aoc25/
│   ├── Main.kt                   # Main entry point
│   ├── framework/
│   │   ├── Day.kt               # Day interface
│   │   ├── Input.kt             # Input processing
│   │   └── DayRegistry.kt       # Day discovery
│   └── days/
│       ├── DayTemplate.kt       # Day template
│       └── Day1.kt              # Your implementations
├── app/src/main/resources/input/
│   ├── 1.input                  # Real inputs
│   └── 1_small.input            # Test inputs
├── scripts/
│   └── new-day.sh               # Day creation script
└── README.md                    # This file
```

## Kotlin Features Used

- **Coroutines & Flows** - Asynchronous processing
- **Sealed Classes** - Type-safe hierarchies
- **Data Classes** - Concise data holders
- **Extension Functions** - Utility functions
- **Reflection** - Automatic day discovery
- **DSL Builders** - Configuration patterns

## Dependencies

- **Kotlin 2.1.0** - Latest Kotlin version
- **Clikt 5.0.1** - Command-line interface
- **Logback 1.5.12** - Logging framework
- **Kotlinx Coroutines 1.9.0** - Asynchronous programming
- **Kotlinx Serialization 1.7.3** - JSON support

## Tips

1. **Use coroutines** for large inputs to avoid memory issues
2. **Enable debug mode** (`--debug`) for detailed logging
3. **Start with small input** (`-s`) for testing
4. **Use flows** for memory-efficient line processing
5. **Leverage extension functions** for clean code

## Example Solutions

Check the `days/` package for examples of different solution patterns and Kotlin idioms.