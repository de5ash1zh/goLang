package main
package main

import (
	"fmt"
)

// Basic struct definition
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Struct with embedded struct
type Address struct {
	Street  string
	City    string
	Country string
}

type Employee struct {
	Person    // Embedded struct
	Address   // Embedded struct
	Company   string
	Salary    float64
}

// Method for Person struct
func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

// Pointer receiver method - can modify the struct
func (p *Person) Birthday() {
	p.Age++
}

// Method for Employee
func (e Employee) AnnualSalary() float64 {
	return e.Salary * 12
}

func main() {
	// Creating a struct
	person1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	fmt.Printf("Person: %+v\n", person1)
	fmt.Printf("Full name: %s\n", person1.FullName())

	// Using pointer receiver method
	person1.Birthday()
	fmt.Printf("After birthday: %+v\n", person1)

	// Creating struct with new
	person2 := new(Person)
	person2.FirstName = "Jane"
	person2.LastName = "Smith"
	person2.Age = 25
	fmt.Printf("Person 2: %+v\n", person2)

	// Creating an employee with embedded structs
	employee := Employee{
		Person: Person{
			FirstName: "Alice",
			LastName:  "Johnson",
			Age:       28,
		},
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
		Company: "Tech Corp",
		Salary:  5000,
	}

	// Accessing fields from embedded structs
	fmt.Printf("\nEmployee: %+v\n", employee)
	fmt.Printf("Employee name: %s\n", employee.FullName())
	fmt.Printf("Employee city: %s\n", employee.City)
	fmt.Printf("Annual salary: $%.2f\n", employee.AnnualSalary())

	// Anonymous struct
	point := struct {
		X, Y int
	}{
		X: 10,
		Y: 20,
	}
	fmt.Printf("\nPoint: %+v\n", point)
}