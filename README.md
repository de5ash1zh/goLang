# Learning GO Lang

## Why Choose Go?

1. Fast Build Time
2. Quick Startup
3. High Performance and Efficiency
4. Excellent Concurrency Model with Goroutines
5. Static Typing and Compilation Benefits

## Project Structure

This repository contains a series of examples demonstrating fundamental Go concepts:

### 1. Hello World (`01_hello/`)

- Basic program structure in Go
- Understanding packages and the `main` function
- Using the `fmt` package for output
- How to run a Go program

### 2. Variables (`02_variables/`)

- Variable declaration and initialization
- Different ways to declare variables:
  - Using `var` keyword
  - Short declaration operator `:=`
- Understanding variable scope
- Basic data types in Go

### 3. Constants (`03_constants/`)

- Constant declaration using `const`
- Understanding immutable values
- Typed and untyped constants
- Working with multiple constants

### 4. Simple Values (`04_simple_values/`)

- Basic data types in Go:
  - Integers
  - Floating-point numbers
  - Strings
  - Booleans
- Type conversion
- Zero values for different types

### 5. Control Structures (`05_for/`)

- For loops in Go
- Different forms of for loops:
  - Traditional for loop
  - For used as while
  - Infinite loops
  - Range-based loops

### 6. Arrays and Slices (`06_arrays_slices/`)

- Fixed-size arrays declaration and usage
- Dynamic slices creation and manipulation
- Slice operations:
  - Appending elements
  - Slicing operations
  - Making slices with capacity
- Copy and reference behavior
- Length vs Capacity

### 7. Maps (`07_maps/`)

- Map creation and initialization
- Key-value operations
- Checking for key existence
- Deleting map entries
- Iterating over maps
- Nested maps
- Map with dynamic keys

### 8. Functions (`08_functions/`)

- Basic function declaration and usage
- Multiple return values
- Named return values
- Variadic functions
- Function as parameters (First-class functions)
- Closures and anonymous functions
- Error handling patterns
- Deferred function calls

### 9. Structs (`09_structs/`)

- Custom type definitions
- Struct declaration and initialization
- Struct methods
- Embedded structs (composition)
- Value vs Pointer receivers
- Anonymous structs
- Field tags and struct annotations
- Struct encapsulation patterns

### 10. Interfaces (`10_interfaces/`)

- Interface declaration and implementation
- Type assertion and type switches
- Empty interface (interface{})
- Multiple interface implementation
- Interface composition
- Common interfaces (Writer, Reader, etc.)
- Interface best practices
- Duck typing in Go

## Basic Go Syntax

```go
package main  // Every Go program starts with package declaration

import "fmt"  // Import required packages

func main() {
    fmt.Println("Hello, World!")  // Entry point of the program
}
```

## Getting Started

To run any example:

1. Navigate to the specific directory
2. Run `go run filename.go`

Example:

```bash
cd 01_hello
go run main.go
```

## Best Practices

1. Always format your code using `go fmt`
2. Use meaningful variable and function names
3. Group related code in packages
4. Follow Go's official style guide
5. Use interfaces for flexibility
6. Handle errors explicitly
7. Prefer composition over inheritance
8. Keep interfaces small and focused

## Next Steps

- Error handling patterns
- Goroutines (Concurrency)
- Channels (Communication)
- Package management
- Testing in Go
- HTTP servers
- Database operations
- JSON handling
- Context usage
- Reflection

## Learning Resources

1. [Go Documentation](https://golang.org/doc/)
2. [Go by Example](https://gobyexample.com/)
3. [Tour of Go](https://tour.golang.org/)
4. [Effective Go](https://golang.org/doc/effective_go)

Feel free to explore each directory for detailed examples and explanations of these concepts. Each example is thoroughly documented with comments explaining the code and concepts.
