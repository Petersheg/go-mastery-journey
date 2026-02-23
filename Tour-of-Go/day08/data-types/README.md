# Go Data Types - Day 08

This module provides an introduction to Go's basic data types and type system concepts, as part of the Tour of Go learning journey.

## Overview

This package (`datatypes`) demonstrates fundamental Go concepts including:
- Basic data type declarations
- Zero values
- Type conversions
- Type inference
- Constants
- Complex numbers

## Basic Data Types in Go

Go has 8 basic data type categories:

1. **bool** - Boolean values (true/false)
2. **string** - Text/string values
3. **int** - Signed integers (int, int8, int16, int32, int64)
4. **uint** - Unsigned integers (uint, uint8, uint16, uint32, uint64, uintptr)
5. **byte** - Alias for uint8
6. **rune** - Alias for int32 (used for Unicode characters)
7. **float** - Floating-point numbers (float32, float64)
8. **complex** - Complex numbers (complex64, complex128)

### Platform Behavior
- `int`, `uint`, and `uintptr` are typically 32 bits wide on 32-bit systems
- On 64-bit systems, they are 64 bits wide
- Use `int` for general integer values unless you have a specific reason to use a sized or unsigned type

## Functions

### `Types()`
Demonstrates variable declaration with different basic types and prints their type and value using the printf format string `%T` (for type) and `%v` (for value).

**Example Variables:**
- `ToBe`: bool
- `Maxint`: uint64
- `z`: complex128

### `ZeroValues()`
Shows Go's zero value concept - the default values assigned to uninitialized variables:
- Numeric types: `0`
- Boolean type: `false`
- String type: `""` (empty string)

Uses `%q` format specifier for quoted string output.

### `TypeConversion()`
Demonstrates explicit type conversion using `T(v)` syntax where `T` is the target type and `v` is the value to convert.

**Example:** Converting between `int` and `float64` for mathematical operations.

```go
var x, y int = 3, 4
var f float64 = math.Sqrt(float64(x*x + y*y))  // Convert int to float64
var z uint = uint(f)                             // Convert float64 to uint
```

### `TypeInference()`
Illustrates Go's type inference system:
- When assigned from typed variables, the new variable takes the same type
- When assigned from untyped numeric constants, the type depends on precision:
  - Integer constants → `int`
  - Floating-point constants → `float64`
  - Complex constants → `complex128`

### `Constant()`
Shows constant declaration and usage. Constants:
- Declared with the `const` keyword
- Can be character, string, boolean, or numeric values
- Cannot use the `:=` syntax
- Cannot be changed after declaration

### Big Number Bitwise Operations
Demonstrates bitwise shifts with large numbers:
- `Big = 1 << 100` - Creates a huge number (2^100)
- `Small = Big >> 99` - Right shifts to get 2

This shows Go's ability to handle arbitrary precision constants.

## Key Concepts

### Type Safety
Go is a statically typed language. Variables have explicit types that must be compatible for operations.

### Zero Values
Every uninitialized variable has a zero value, preventing undefined behavior:
```go
var i int        // 0
var f float32    // 0
var b bool       // false
var s string     // ""
```

### Type Conversion vs Type Assertion
- **Type Conversion**: `T(v)` converts value v to type T
- Must be explicit - Go has no implicit type conversion

### Format Specifiers Used
- `%T` - Type of the value
- `%v` - Value in default format
- `%q` - Quoted string format
- `%i` - Integer format

## Package Dependencies

```go
import (
	"fmt"           // Formatted I/O
	"math"          // Mathematical functions
	"math/cmplx"    // Complex number operations
)
```

## Usage

This module is part of the Tour of Go learning path and is meant to be studied interactively as you learn Go fundamentals.

## Related Topics

- Variable declaration and initialization
- Type system in Go
- Constants and iota
- Complex numbers and mathematical operations
- Type assertions (covered in later topics)
