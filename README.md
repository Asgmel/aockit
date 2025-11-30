# Advent of Code Toolkit

A reusable Go toolkit for solving [Advent of Code](https://adventofcode.com/) puzzles. This library provides common utilities, data structures, and helper functions to streamline your AoC solutions.

## Installation

```bash
go get github.com/asgmel/aockit
```

## Features

### 📥 Input Reading (`input` package)

Simplify reading and parsing puzzle inputs with multiple convenience functions:

- **`ReadInputString(path)`** - Read entire file as a single string
- **`ReadInputLines(path)`** - Split input into lines
- **`ReadInputLetters(path)`** - Parse input into a 2D grid of individual characters
- **`ReadInputRegex(path, regex)`** - Extract all matches for a regex pattern

#### Example
```go
import "github.com/asgmel/aockit/input"

// Read all lines
lines := input.ReadInputLines("input.txt")

// Parse as character grid
grid := input.ReadInputLetters("input.txt")

// Extract numbers
numbers := input.ReadInputRegex("input.txt", `\d+`)
```

### 📍 Position Helpers (`grid` package)

Work with 2D coordinates effortlessly:

- **`Position{X, Y}`** - Simple 2D coordinate struct
- **`ToString()`** - Convert position to string representation
- **`WithinBounds(matrix)`** - Check if position is within grid bounds

#### Example
```go
import "github.com/asgmel/aockit/grid"

pos := grid.Position{X: 5, Y: 10}
if pos.WithinBounds(grid) {
    // Position is valid
}
```

### 🛠️ Utility Functions (`utils` package)

Common operations you'll use repeatedly:

- **`AbsoluteValue(int)`** - Get absolute value
- **`SumIntSlice([]int)`** - Sum all integers in a slice
- **`ConvertStrSliceToIntSlice([]string)`** - Convert string slice to integers
- **`CopySlice[T]([]T)`** - Create a deep copy of any slice
- **`FilterDuplicates[T]([]T)`** - Remove duplicate values

#### Example
```go
import "github.com/asgmel/aockit/utils"

sum := utils.SumIntSlice([]int{1, 2, 3, 4, 5}) // 15
unique := utils.FilterDuplicates([]string{"a", "b", "a"}) // ["a", "b"]
```

## Usage

Import the packages you need in your Advent of Code solutions:

```go
package main

import (
    "fmt"
    "github.com/asgmel/aockit/input"
    "github.com/asgmel/aockit/utils"
)

func main() {
    lines := input.ReadInputLines("day01/input.txt")
    numbers := utils.ConvertStrSliceToIntSlice(lines)
    total := utils.SumIntSlice(numbers)
    fmt.Println(total)
}
```

## Requirements

- Go 1.25.4 or higher

## License

MIT