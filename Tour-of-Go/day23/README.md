# Day 23 - Interfaces Continued

Covers three interface concepts: nil underlying values, nil interface values, and empty interfaces.

## Concepts

**Interface values with nil underlying values**
- An interface can hold a nil concrete value and still be non-nil itself
- Methods can gracefully handle a nil receiver without panicking

**Nil interface values**
- A nil interface holds neither a value nor a concrete type
- Calling a method on a nil interface causes a runtime panic

**Empty interface (`interface{}`)**
- Specifies zero methods — every type satisfies it
- Used to handle values of unknown type (e.g. `fmt.Print`)

## Examples

```go
// Nil underlying value — safe, M() handles nil receiver
var t *T
i = t
i.M() // prints <nil>

// Nil interface — panics at runtime
var i I
i.M() // panic: nil pointer dereference

// Empty interface — holds any type
var i interface{}
i = 42
i = "Hello"
```
