# Day 25 - Errors

## Concepts
- The `error` built-in interface with a single `Error() string` method
- Returning custom errors from functions
- Checking `err != nil` to handle failure vs success

## Custom Error Types
Implement the `error` interface on a struct to carry context:

```go
type MyError struct {
    When time.Time
    What string
}

func (e MyError) Error() string {
    return fmt.Sprintf("at %v, %s", e.What, e.When)
}
```

## Exercise: ErrNegativeSqrt
Defined `ErrNegativeSqrt` as a `float64` type implementing `error`, used inside `Sqrt` to signal invalid input (negative numbers).
