# Day 08 - Data Types & For Loops

## Modules

### data-types/
Demonstrates Go's basic data type system including:
- Basic type declarations (bool, int, uint, float, complex)
- Zero values
- Type conversions
- Type inference
- Constants

### for/
Explores Go's looping construct with various implementations:
- **SumUp()** - Standard for loop with init, condition, and post statements
- **SumUpWithoutInit()** - For loop with optional init and post statements
- **ForAsWhile()** - Using for loop as a while loop (condition only)
- **Forever()** - Infinite loop pattern

## Key Takeaways

- Go has only **one looping construct**: the for loop
- Unlike C/Java, for loops don't need parentheses around components
- Braces `{ }` are always required
- For loops can be used as while loops by omitting init/post statements
- Infinite loops are created with just `for {}`
