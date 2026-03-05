# Day 16: Advanced Slice Operations in Go

This directory demonstrates more advanced slice operations in Go, based on what was learned so far:

## Covered Topics

- **Creating Slices with `make`**
  - Use `make([]T, len)` or `make([]T, len, cap)` to create slices with specified length and capacity.
- **Slices of Slices**
  - Slices can contain other slices, enabling structures like 2D arrays (e.g., Tic Tac Toe board).
- **Appending to Slices**
  - Use the built-in `append` function to add elements to a slice. If needed, Go automatically allocates a new underlying array.

## Files
- `main.go`: Entry point for running slice examples.
- `slice-cond.go`: Contains all code and explanations for the above slice operations.

> **Note:** This README covers only the topics and code present so far.
