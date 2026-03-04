# Day 15: Arrays and Slices in Go

This directory contains Go code examples and explanations for arrays and slices, following the Tour of Go. The focus so far has been on slices, covering the following topics:

## Covered Topics

- **What is a Slice?**
	- Slices are dynamically-sized, flexible views into arrays.
	- Syntax: `a[low:high]` creates a slice from an array.
- **Slices as References**
	- Slices reference the underlying array, so changes to a slice affect the array and other slices referencing it.
- **Slice Literals**
	- Slice literals are like array literals but without a length: `[]int{1, 2, 3}`.
- **Slice Defaults**
	- Omitting low/high bounds uses defaults: `a[:3]`, `a[2:]`, `a[:]`.
- **Slice Length and Capacity**
	- Use `len(s)` and `cap(s)` to get a slice's length and capacity.
	- Slices can be re-sliced within their capacity.
- **Nil Slices**
	- The zero value of a slice is `nil`, with length and capacity 0.

## Files
- `main.go`: Entry point, calls slice demonstration functions.
- `slice.go`: Contains all slice-related code and explanations.

> **Note:** Only slices are covered so far. More on arrays and advanced slice operations will be added later.
