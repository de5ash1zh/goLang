package main
package main

import (
	"fmt"
)

// Basic function
func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Named return values
func rectangle(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // naked return
}

// Variadic function
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function as parameter
func calculate(operation func(int, int) int, a, b int) int {
	return operation(a, b)
}

// Closure (anonymous function)
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	// Basic function call
	message := greet("Gopher")
	fmt.Println(message)

	// Multiple return values
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", result)
	}

	// Division by zero example
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Named return values
	area, perimeter := rectangle(5, 3)
	fmt.Printf("Rectangle - Area: %.2f, Perimeter: %.2f\n", area, perimeter)

	// Variadic function
	fmt.Printf("Sum of 1, 2, 3: %d\n", sum(1, 2, 3))
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum of slice: %d\n", sum(numbers...))

	// Function as parameter
	add := func(x, y int) int { return x + y }
	multiply := func(x, y int) int { return x * y }

	fmt.Printf("Calculate (add): %d\n", calculate(add, 5, 3))
	fmt.Printf("Calculate (multiply): %d\n", calculate(multiply, 5, 3))

	// Closure example
	incrementCounter := counter()
	fmt.Printf("Counter: %d\n", incrementCounter())
	fmt.Printf("Counter: %d\n", incrementCounter())
	fmt.Printf("Counter: %d\n", incrementCounter())
}