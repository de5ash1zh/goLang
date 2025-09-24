package main
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Example 1: Struct with tags
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Password  string    `json:"-"` // Will be omitted from JSON
}

// Example 2: Custom JSON marshaling
type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("invalid duration")
	}
}

// Example 3: Builder Pattern
type Computer struct {
	CPU       string
	RAM       int
	Storage   int
	GPU       string
	Monitor   bool
	Keyboard  bool
	Mouse     bool
}

type ComputerBuilder struct {
	computer *Computer
}

func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{computer: &Computer{}}
}

func (b *ComputerBuilder) WithCPU(cpu string) *ComputerBuilder {
	b.computer.CPU = cpu
	return b
}

func (b *ComputerBuilder) WithRAM(ram int) *ComputerBuilder {
	b.computer.RAM = ram
	return b
}

func (b *ComputerBuilder) WithStorage(storage int) *ComputerBuilder {
	b.computer.Storage = storage
	return b
}

func (b *ComputerBuilder) WithGPU(gpu string) *ComputerBuilder {
	b.computer.GPU = gpu
	return b
}

func (b *ComputerBuilder) WithPeripherals(monitor, keyboard, mouse bool) *ComputerBuilder {
	b.computer.Monitor = monitor
	b.computer.Keyboard = keyboard
	b.computer.Mouse = mouse
	return b
}

func (b *ComputerBuilder) Build() *Computer {
	return b.computer
}

// Example 4: Inheritance-like behavior with embedding
type Animal struct {
	Name   string
	Species string
}

func (a *Animal) Describe() string {
	return fmt.Sprintf("%s is a %s", a.Name, a.Species)
}

type Pet struct {
	*Animal
	Owner  string
	HomeAddress string
}

func (p *Pet) FullInfo() string {
	return fmt.Sprintf("%s and belongs to %s", p.Describe(), p.Owner)
}

// Example 5: Method chaining
type StringBuilder struct {
	str string
}

func (sb *StringBuilder) Append(s string) *StringBuilder {
	sb.str += s
	return sb
}

func (sb *StringBuilder) AppendLine(s string) *StringBuilder {
	sb.str += s + "\n"
	return sb
}

func (sb *StringBuilder) ToString() string {
	return sb.str
}

func main() {
	// Example 1: JSON marshaling with struct tags
	user := User{
		ID:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
		Password:  "secret123",
	}
	
	jsonData, _ := json.MarshalIndent(user, "", "  ")
	fmt.Printf("JSON with struct tags:\n%s\n\n", jsonData)

	// Example 2: Custom JSON marshaling
	config := struct {
		Timeout Duration `json:"timeout"`
	}{
		Timeout: Duration{time.Hour},
	}
	
	jsonConfig, _ := json.MarshalIndent(config, "", "  ")
	fmt.Printf("Custom marshaled JSON:\n%s\n\n", jsonConfig)

	// Example 3: Builder Pattern
	computer := NewComputerBuilder().
		WithCPU("Intel i7").
		WithRAM(32).
		WithStorage(1000).
		WithGPU("NVIDIA RTX 3080").
		WithPeripherals(true, true, true).
		Build()
	
	fmt.Printf("Built computer: %+v\n\n", computer)

	// Example 4: Inheritance-like behavior
	pet := &Pet{
		Animal: &Animal{
			Name:    "Max",
			Species: "Dog",
		},
		Owner:       "Alice",
		HomeAddress: "123 Pet Street",
	}
	
	fmt.Println(pet.FullInfo())

	// Example 5: Method chaining
	sb := &StringBuilder{}
	result := sb.Append("Hello").
		AppendLine(" World!").
		Append("How").
		AppendLine(" are you?").
		ToString()
	
	fmt.Printf("\nMethod chaining result:\n%s", result)
}