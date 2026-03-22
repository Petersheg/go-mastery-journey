# Day 22 - Interfaces

## Concepts Covered

- **Interface definition** – a set of method signatures; any type implementing those methods satisfies the interface
- **Implicit implementation** – no `implements` keyword; a type satisfies an interface simply by implementing its methods
- **Pointer vs value receivers** – only `*Vertex` implements `Abser` because `Abs()` is defined on the pointer receiver
- **Interface values** – internally a tuple of `(value, type)`; calling a method dispatches to the underlying concrete type

## Key Examples

| Function | Description |
|---|---|
| `Interface()` | Assigns `MyFloat` and `*Vertex` to an `Abser` interface |
| `ExplicitInterface()` | Shows implicit implementation — `T` satisfies `I` without any declaration |
| `InterfaceValue()` | Demonstrates interface value internals using `describe()` to print `(value, type)` |

## Notes

- Assigning a value type (`Vertex`) to an interface that only the pointer type (`*Vertex`) implements causes a compile error
- Implicit interfaces decouple interface definitions from implementations, enabling use across packages without prearrangement
