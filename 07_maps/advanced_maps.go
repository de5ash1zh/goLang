package main
package main

import (
	"fmt"
	"sort"
	"sync"
)

// Custom type for map value
type UserInfo struct {
	Age     int
	City    string
	Hobbies []string
}

// SafeMap provides a concurrent-safe map
type SafeMap struct {
	sync.RWMutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.Lock()
	defer sm.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.RLock()
	defer sm.RUnlock()
	val, ok := sm.data[key]
	return val, ok
}

func main() {
	// Example 1: Map with complex value type
	users := map[string]UserInfo{
		"john": {
			Age:     28,
			City:    "New York",
			Hobbies: []string{"reading", "hiking"},
		},
		"alice": {
			Age:     25,
			City:    "London",
			Hobbies: []string{"painting", "music"},
		},
	}

	// Accessing complex map values
	fmt.Println("John's info:", users["john"])
	fmt.Println("Alice's hobbies:", users["alice"].Hobbies)

	// Example 2: Sorting map keys
	scores := map[string]int{
		"Eve":   95,
		"Bob":   87,
		"Alice": 92,
		"David": 89,
	}

	// Get all keys
	var keys []string
	for k := range scores {
		keys = append(keys, k)
	}

	// Sort keys
	sort.Strings(keys)

	// Print scores in sorted order
	fmt.Println("\nScores in alphabetical order:")
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, scores[k])
	}

	// Example 3: Map as a set
	seen := make(map[string]bool)
	words := []string{"apple", "banana", "apple", "cherry", "banana", "date"}

	fmt.Println("\nUnique words:")
	for _, word := range words {
		if !seen[word] {
			seen[word] = true
			fmt.Println(word)
		}
	}

	// Example 4: Concurrent-safe map
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	// Concurrent writes
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			safeMap.Set(fmt.Sprintf("key%d", n), n)
		}(i)
	}

	wg.Wait()

	// Print safe map contents
	fmt.Println("\nSafe map contents:")
	for i := 0; i < 10; i++ {
		if val, ok := safeMap.Get(fmt.Sprintf("key%d", i)); ok {
			fmt.Printf("key%d: %d\n", i, val)
		}
	}

	// Example 5: Map of functions
	operations := map[string]func(int, int) int{
		"add":      func(a, b int) int { return a + b },
		"subtract": func(a, b int) int { return a - b },
		"multiply": func(a, b int) int { return a * b },
		"divide":   func(a, b int) int { if b != 0 { return a / b }; return 0 },
	}

	// Using function map
	fmt.Println("\nFunction map results:")
	a, b := 15, 5
	for op, fn := range operations {
		fmt.Printf("%s: %d\n", op, fn(a, b))
	}
}