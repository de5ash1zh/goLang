# Learning GO Lang - A Practical Guide

## Why Go?

- **Speed**: Fast compilation and runtime performance
- **Simplicity**: Clean syntax and standard formatting
- **Safety**: Strong static typing and built-in concurrency
- **Modern**: Built for multicore machines and networked services
- **Batteries Included**: Rich standard library

## Quick Start

```bash
# Run any example
cd 01_hello
go run main.go
```

## Core Concepts with Examples

### 1. Basic Structure (`01_hello/`)

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Every Go program starts with:

- `package` declaration (package `main` for executables)
- `import` statements for required packages
- `main()` function as the entry point

### 2. Variables (`02_variables/`)

```go
// Different ways to declare variables
var name string = "John"     // Explicit type
age := 25                    // Type inference
var isActive = true         // Type inference with var
var count int               // Declaration with zero value (0)

// Multiple declarations
var (
    firstName = "John"
    lastName  = "Doe"
    height    = 175.5
)
```

### 3. Constants (`03_constants/`)

```go
// Typed and untyped constants
const Pi = 3.14159          // Untyped float
const MaxAge int = 120      // Typed int

// Multiple constants
const (
    Sunday = iota          // 0
    Monday                 // 1
    Tuesday                // 2
)
```

### 4. Simple Values (`04_simple_values/`)

```go
// Basic types
str := "hello"              // string
num := 42                   // int
pi := 3.14                 // float64
isTrue := true             // boolean

// Type conversion
height := 1.75             // float64
heightInt := int(height)   // Convert to int (1)
```

### 5. Control Structures (`05_for/`)

````go
// Standard for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// For as while
sum := 0
for sum < 100 {
    sum += 10
}

// Range-based loop
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("numbers[%d] = %d\n", index, value)
}

### 6. Arrays and Slices (`06_arrays_slices/`)
```go
// Arrays (fixed size)
var numbers [5]int                     // Array of 5 integers
colors := [3]string{"red", "green", "blue"}

// Slices (dynamic size)
names := []string{"John", "Jane"}      // Slice literal
numbers := make([]int, 3, 5)          // len=3, cap=5

// Slice operations
numbers = append(numbers, 4, 5)        // Add elements
slice := numbers[1:4]                  // Slicing
````

### 7. Maps (`07_maps/`)

```go
// Create and initialize maps
ages := map[string]int{
    "John": 25,
    "Jane": 30,
}

// Map operations
ages["Bob"] = 35                       // Add/Update
age, exists := ages["John"]            // Safe access
delete(ages, "Jane")                   // Delete

// Iterate over map
for name, age := range ages {
    fmt.Printf("%s is %d years old\n", name, age)
}
```

### 8. Functions (`08_functions/`)

```go
// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Variadic function
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

// Function closure
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

### 9. Structs (`09_structs/`)

```go
// Define and use structs
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// Struct method
func (p Person) Greet() string {
    return fmt.Sprintf("Hi, I'm %s", p.Name)
}

// Create and use struct
person := Person{Name: "John", Age: 25}
fmt.Println(person.Greet())

// Embedded structs
type Employee struct {
    Person            // Embedding
    Department string
}
```

### 10. Interfaces (`10_interfaces/`)

```go
// Define interface
type Shape interface {
    Area() float64
}

// Implement interface
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// Use interface
func PrintArea(s Shape) {
    fmt.Printf("Area: %0.2f\n", s.Area())
}
```

## Best Practices

1. **Code Organization**

   ```go
   // Group related functions
   package users

   // Use meaningful names
   func CreateUser(name string) (*User, error)
   func DeleteUser(id string) error
   ```

2. **Error Handling**

   ```go
   if err != nil {
       return fmt.Errorf("failed to process: %v", err)
   }
   ```

3. **Interface Design**
   ```go
   // Keep interfaces small
   type Reader interface {
       Read(p []byte) (n int, err error)
   }
   ```

## Learning Path

1. Start with basic syntax (`01_hello` → `05_for`)
2. Move to data structures (`06_arrays_slices` → `07_maps`)
3. Learn functions and error handling (`08_functions`)
4. Master structs and interfaces (`09_structs` → `10_interfaces`)
5. Explore advanced examples in each section

## Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go)

Feel free to explore each directory for detailed examples. Each file is thoroughly documented with comments explaining the concepts.

## Learning Resources

1. [Go Documentation](https://golang.org/doc/)
2. [Go by Example](https://gobyexample.com/)
3. [Tour of Go](https://tour.golang.org/)
4. [Effective Go](https://golang.org/doc/effective_go)

Feel free to explore each directory for detailed examples and explanations of these concepts. Each example is thoroughly documented with comments explaining the code and concepts.
