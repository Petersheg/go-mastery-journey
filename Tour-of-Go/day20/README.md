# Day 20 - Methods

Learned about methods in Go - functions with receiver arguments that attach to types.

## What I Learned
- Methods are functions with a special receiver argument between `func` and the method name
- Can define methods on struct types (e.g., `Vertex`)
- Can define methods on custom non-struct types (e.g., `MyFloat`)
- Methods vs regular functions - same functionality, different syntax
- Cannot declare methods on types from other packages (including built-in types)

## Examples
- **Receiver Method**: `Abs()` method on `Vertex` struct calculating distance from origin
- **Ordinary Function**: Same `Abs` logic as a regular function for comparison
- **Non-Struct Method**: `Abs()` method on custom `MyFloat` type

## Run
```sh
go run main.go methods.go
```