package main
package main

import (
	"fmt"
)

func main() {
	// Creating a map
	fmt.Println("Maps in Go:")
	
	// Method 1: Using make
	scores := make(map[string]int)
	scores["Alice"] = 98
	scores["Bob"] = 87
	scores["Charlie"] = 92
	fmt.Printf("Student scores: %v\n", scores)

	// Method 2: Map literal
	countries := map[string]string{
		"US": "United States",
		"UK": "United Kingdom",
		"IN": "India",
		"FR": "France",
	}
	fmt.Printf("Countries: %v\n", countries)

	// Accessing values
	aliceScore := scores["Alice"]
	fmt.Printf("Alice's score: %d\n", aliceScore)

	// Checking if key exists
	if score, exists := scores["David"]; exists {
		fmt.Printf("David's score: %d\n", score)
	} else {
		fmt.Println("David's score not found")
	}

	// Deleting from map
	delete(scores, "Bob")
	fmt.Printf("After deleting Bob's score: %v\n", scores)

	// Length of map
	fmt.Printf("Number of scores: %d\n", len(scores))

	// Iterating over map
	fmt.Println("\nIterating over countries:")
	for code, name := range countries {
		fmt.Printf("%s: %s\n", code, name)
	}

	// Nested maps
	cities := map[string]map[string]string{
		"USA": {
			"capital": "Washington D.C.",
			"largest": "New York City",
		},
		"India": {
			"capital": "New Delhi",
			"largest": "Mumbai",
		},
	}
	fmt.Printf("\nNested map - Cities: %v\n", cities)
}