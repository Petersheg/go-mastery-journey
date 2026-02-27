# Day 11 - Switch Statements

Explored Go's switch statement as a cleaner alternative to if-else chains.

## Key Concepts

- **Switch with condition**: Evaluates cases against a condition expression
- **Switch without condition**: Same as `switch true`, useful for long if-else chains
- **Automatic break**: Go only runs the selected case, no fallthrough by default
- **Runtime info**: Used `runtime` package to get system information (OS, architecture, version, CPU count)

## Examples

- `SwitchOS()`: Detects operating system
- `SwitchDay()`: Calculates days until Saturday
- `SwitchTrue()`: Time-based greeting using switch without condition
