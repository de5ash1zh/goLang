package main
package main

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

// Example 1: Interface composition
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

// Example 2: Stringer interface
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

// Example 3: Custom error interface
type NetworkError struct {
	Code    int
	Message string
}

func (e NetworkError) Error() string {
	return fmt.Sprintf("network error: %s (code: %d)", e.Message, e.Code)
}

// Example 4: Sort interface implementation
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// Example 5: Custom ReadWriter implementation
type StringReadWriter struct {
	content string
	pos     int
}

func NewStringReadWriter(initial string) *StringReadWriter {
	return &StringReadWriter{content: initial}
}

func (s *StringReadWriter) Read(p []byte) (n int, err error) {
	if s.pos >= len(s.content) {
		return 0, io.EOF
	}

	n = copy(p, []byte(s.content[s.pos:]))
	s.pos += n
	return n, nil
}

func (s *StringReadWriter) Write(p []byte) (n int, err error) {
	s.content += string(p)
	return len(p), nil
}

func (s *StringReadWriter) String() string {
	return s.content
}

// Example 6: Interface type assertion and type switches
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

// Function that uses type assertions
func describe(i interface{}) {
	switch v := i.(type) {
	case Shape:
		fmt.Printf("Shape with area: %.2f\n", v.Area())
	case string:
		fmt.Printf("String with length: %d\n", len(v))
	case int:
		fmt.Printf("Integer with value: %d\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	// Example 1: Interface composition
	rw := NewStringReadWriter("Hello")
	fmt.Fprintf(rw, ", World!") // Uses Write
	
	buf := make([]byte, 100)
	n, _ := rw.Read(buf) // Uses Read
	fmt.Printf("Read %d bytes: %s\n\n", n, buf[:n])

	// Example 2: Stringer interface
	person := Person{"Alice", 30}
	fmt.Printf("Person: %s\n\n", person) // Uses String() method

	// Example 3: Custom error
	err := NetworkError{404, "Page not found"}
	fmt.Printf("Error: %v\n\n", err) // Uses Error() method

	// Example 4: Sort interface
	people := []Person{
		{"Bob", 31},
		{"Alice", 25},
		{"Charlie", 35},
	}
	
	sort.Sort(ByAge(people))
	fmt.Println("Sorted by age:")
	for _, p := range people {
		fmt.Printf("  %s\n", p)
	}
	fmt.Println()

	// Example 5: Interface type assertion
	var shape Shape = Circle{Radius: 5}
	
	if circle, ok := shape.(Circle); ok {
		fmt.Printf("Circle with radius: %.2f\n", circle.Radius)
	}

	// Example 6: Type switches
	fmt.Println("\nType switch examples:")
	describe(Square{Side: 4})
	describe("Hello")
	describe(42)
	describe(3.14)

	// Example 7: Empty interface as type constraint
	data := []interface{}{
		42,
		"Hello",
		Person{"Bob", 25},
		Circle{5},
	}

	fmt.Println("\nProcessing various types:")
	for _, item := range data {
		describe(item)
	}

	// Example 8: Interface composition with embedding
	type LogWriter interface {
		io.Writer
		Log(message string)
	}

	type ConsoleLogger struct {
		prefix string
	}

	func (c ConsoleLogger) Write(p []byte) (n int, err error) {
		fmt.Printf("%s: %s", c.prefix, string(p))
		return len(p), nil
	}

	func (c ConsoleLogger) Log(message string) {
		c.Write([]byte(message + "\n"))
	}

	logger := ConsoleLogger{prefix: "DEBUG"}
	logger.Log("This is a debug message")
}