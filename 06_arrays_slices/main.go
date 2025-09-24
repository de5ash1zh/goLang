package main

import (
	"fmt"
)

func main() {
	// Arrays
	fmt.Println("Arrays in Go:")
	// Fixed-size array
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Printf("Array: %v\n", numbers)

	// Array literal
	fruits := [3]string{"apple", "banana", "orange"}
	fmt.Printf("Fruits: %v\n", fruits)

	// Slices
	fmt.Println("\nSlices in Go:")
	// Create slice from array
	numberSlice := numbers[1:4]
	fmt.Printf("Slice from array: %v\n", numberSlice)

	// Create slice with make
	dynamicSlice := make([]int, 3, 5)
	fmt.Printf("Slice length: %d, capacity: %d\n", len(dynamicSlice), cap(dynamicSlice))

	// Append to slice
	dynamicSlice = append(dynamicSlice, 1, 2, 3)
	fmt.Printf("After append: %v\n", dynamicSlice)

	// Slice literal
	cities := []string{"New York", "London", "Tokyo"}
	fmt.Printf("Cities: %v\n", cities)

	// Copy slices
	destination := make([]string, len(cities))
	copied := copy(destination, cities)
	fmt.Printf("Copied %d elements: %v\n", copied, destination)

	// Slice with variadic capacity
	letters := []string{}
	fmt.Printf("Empty slice - Length: %d, Capacity: %d\n", len(letters), cap(letters))
	letters = append(letters, "a", "b", "c")
	fmt.Printf("After append - Length: %d, Capacity: %d\n", len(letters), cap(letters))
}