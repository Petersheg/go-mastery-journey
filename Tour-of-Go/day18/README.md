# Day 18 – Maps in Go

This section covers the basics of Go maps:

- Declaring and initializing maps
- Using map literals
- Mutating maps (adding, updating, deleting elements)
- Practical examples with custom structs as values

See `main.go` and `maps.go` for code demonstrations.

**Key Concepts:**
- Maps are Go's built-in associative data structure.
- The zero value of a map is `nil`.
- Use `make()` to initialize a map before use.
- Map literals allow concise initialization.
- Maps can be mutated by adding, updating, or deleting keys.

---

**Example:**
```go
m := make(map[string]int)
m["answer"] = 42
fmt.Println(m["answer"])
```

For more, see the code files in this folder.
