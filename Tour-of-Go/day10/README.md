# Day 10 - Loops and Functions

Implementation of square root calculation using Newton's method, combining loops and if statements.

## Exercise

Compute the square root of a number using iterative approximation:
- Formula: `z -= (z*z - x) / (2*z)`
- Stops when change is less than epsilon (0.0000001)
- Maximum 10 iterations

## Usage

```go
result := Sqrt(2.0)
```

Compares results with `math.Sqrt` from the standard library.
