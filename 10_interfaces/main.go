package main
package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle implements Shape
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Writer interface
type Writer interface {
	Write(string)
}

// ConsoleWriter implements Writer
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(s string) {
	fmt.Println("Console:", s)
}

// FileWriter implements Writer
type FileWriter struct {
	FilePath string
}

func (fw FileWriter) Write(s string) {
	fmt.Printf("Writing to file %s: %s\n", fw.FilePath, s)
}

// Function that accepts interface
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func WriteMessage(w Writer, message string) {
	w.Write(message)
}

// Empty interface example
func PrintAnything(v interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", v, v)
}

func main() {
	// Using shapes
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	fmt.Println("Circle:")
	PrintShapeInfo(circle)

	fmt.Println("\nRectangle:")
	PrintShapeInfo(rectangle)

	// Slice of interfaces
	shapes := []Shape{circle, rectangle}
	fmt.Println("\nAll shapes:")
	for _, shape := range shapes {
		PrintShapeInfo(shape)
		fmt.Println()
	}

	// Using writers
	console := ConsoleWriter{}
	file := FileWriter{FilePath: "output.txt"}

	WriteMessage(console, "Hello from console!")
	WriteMessage(file, "Hello from file!")

	// Empty interface example
	fmt.Println("\nEmpty interface examples:")
	PrintAnything(42)
	PrintAnything("Hello")
	PrintAnything(true)
	PrintAnything(circle)

	// Type assertion
	var i interface{} = "Hello, Go!"
	if str, ok := i.(string); ok {
		fmt.Printf("\nValue is string: %s\n", str)
	}

	// Type switch
	var j interface{} = 42
	switch v := j.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}