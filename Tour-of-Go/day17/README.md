# Day 17 – Slices Exercise

This folder contains a solution to the "Slices" exercise from the Go Tour.

## Exercise Description
Implement a function `Pic(dx, dy int) [][]uint8` that returns a slice of length `dy`, where each element is a slice of `dx` 8-bit unsigned integers. The values can be generated using interesting functions such as `(x+y)/2`, `x*y`, or `x^y`.

## Files
- `excercise.go`: Contains the implementation of the `Pic` function and supporting code.

## How to Run
1. Ensure you have Go installed.
2. Run the file:
   ```sh
   go run excercise.go
   ```

## Notes
- The implementation uses random grayscale values for demonstration.
- You can modify the value generation logic in `getRand` for different patterns.
