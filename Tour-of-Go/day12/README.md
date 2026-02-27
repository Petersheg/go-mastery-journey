# Day 12 - Defer Statements

Learned about defer statements for delaying function execution until the surrounding function returns.

## Key Concepts

- **Defer execution**: Deferred function runs when surrounding function returns
- **Immediate evaluation**: Arguments are evaluated immediately, but execution is deferred
- **LIFO order**: Deferred calls are stacked and executed in last-in-first-out order

## Examples

- `Defer()`: Basic defer usage - prints "Hello" then "World"
- `StackingDefer()`: Demonstrates LIFO order with deferred calls in a loop (counts down from 9 to 0)
