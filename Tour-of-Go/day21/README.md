# Day 21 - Pointer Receivers

## Concepts Covered

- Pointer receivers (`*T`) vs value receivers — pointer receivers can modify the original value
- Rewriting methods as plain functions to understand the difference
- Pointer indirection: Go auto-dereferences `v.Scale(5)` to `(&v).Scale(5)` for pointer receiver methods
- Reverse indirection: `p.Abs()` is interpreted as `(*p).Abs()` for value receiver methods
- Choosing receivers: prefer pointer receivers to mutate state or avoid copying large structs; keep receiver types consistent across a type's methods

## Files

- `pointer_receiver.go` — Vertex examples demonstrating all the above concepts via `Calc` through `Calc5`
- `main.go` — entry point calling the calc functions
