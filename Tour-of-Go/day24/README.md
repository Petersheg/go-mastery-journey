# Day 24 - Interfaces Continued: Type Assertions, Type Switches, and Stringers

## Topics Covered

- Type Assertions: access an interface's underlying concrete value with `i.(T)`, with optional `ok` check to avoid panics
- Type Switches: series of type assertions using `switch v := i.(type)` to branch on concrete type
- Stringers: implementing the `fmt.Stringer` interface (`String() string`) so types can describe themselves as strings

## Exercise

Implemented `fmt.Stringer` on `IPAddr [4]byte` to print IP addresses in dotted-quad format (e.g. `1.2.3.4`).
