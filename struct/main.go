package main

import "fmt"



type user struct {
	name  string
	email string
}

// ============================================
// ADDITION 1: Nested Struct
// ============================================

type address struct {
	city    string
	country string
}

type employee struct {
	name    string
	email   string
	workAdd address // nested struct
}

// ============================================
// ADDITION 2: Struct Embedding (Composition)
// ============================================

type manager struct {
	user   // embedded struct (no field name)
	team   int
	branch string
}

// ============================================
// ADDITION 3: Method with Pointer Receiver
// ============================================

type product struct {
	name  string
	price float64
	stock int
}

// Method to display product info
func (p product) display() {
	fmt.Printf("  %s: $%.2f (Stock: %d)\n", p.name, p.price, p.stock)
}

// Method to apply discount (needs pointer to modify)
func (p *product) applyDiscount(percent float64) {
	p.price = p.price * (1 - percent/100)
}

// Method to reduce stock
func (p *product) sell(quantity int) bool {
	if p.stock >= quantity {
		p.stock -= quantity
		return true
	}
	return false
}

// ============================================
// ADDITION 4: Constructor Pattern
// ============================================

type account struct {
	owner   string
	balance float64
}

// Constructor function - Go convention: NewXxx
func NewAccount(owner string, initialBalance float64) *account {
	return &account{
		owner:   owner,
		balance: initialBalance,
	}
}

// ============================================
// ADDITION 5: Pointer to Struct (Accessing Fields)
// ============================================

type config struct {
	host    string
	port    int
	verbose bool
}

func updateConfig(c *config, host string, port int) {
	c.host = host
	c.port = port
}

func main() {
	

	sabuj := user{
		name:  "Sabuj",
		email: "sabuj@gmail.com",
	}

	fmt.Printf("%+v\n", sabuj)
	fmt.Println(sabuj.email)
	fmt.Println()

	// ============================================
	// ADDITION 1: Nested Struct
	// ============================================
	fmt.Println("=== NESTED STRUCT ===")
 emp := employee{
		name:    "Rahim",
		email:   "rahim@company.com",
		workAdd: address{city: "Dhaka", country: "Bangladesh"},
	}

	fmt.Printf("Name: %s\n", emp.name)
	fmt.Printf("City: %s\n", emp.workAdd.city)
	fmt.Printf("Country: %s\n", emp.workAdd.country)
	fmt.Println()

	// ============================================
	// ADDITION 2: Struct Embedding (Promoted Fields)
	// ============================================
	fmt.Println("=== STRUCT EMBEDDING ===")
	mgr := manager{
		user:   user{name: "Karim", email: "karim@company.com"},
		team:   15,
		branch: "Dhaka",
	}

	// Access embedded struct fields directly (promoted)
	fmt.Printf("Name: %s\n", mgr.name)       // promoted from user
	fmt.Printf("Email: %s\n", mgr.email)     // promoted from user
	fmt.Printf("Team: %d\n", mgr.team)
	fmt.Printf("Branch: %s\n", mgr.branch)
	fmt.Println()

	// ============================================
	// ADDITION 3: Methods on Struct
	// ============================================
	fmt.Println("=== METHODS ON STRUCT ===")
	p1 := product{name: "Laptop", price: 999.99, stock: 10}
	p2 := product{name: "Mouse", price: 29.99, stock: 50}

	p1.display()
	p2.display()

	// Apply discount (modifies price)
	p1.applyDiscount(10) // 10% off
	fmt.Println("\nAfter 10% discount:")
	p1.display()

	// Sell product
	if p2.sell(5) {
		fmt.Println("\nSold 5 mice!")
	}
	p2.display()
	fmt.Println()

	// ============================================
	// ADDITION 4: Constructor Pattern
	// ============================================
	fmt.Println("=== CONSTRUCTOR PATTERN ===")
	acc := NewAccount("Sabuj", 1000.00)
	fmt.Printf("Owner: %s\n", acc.owner)
	fmt.Printf("Balance: $%.2f\n", acc.balance)
	fmt.Println()

	// ============================================
	// ADDITION 5: Pointer to Struct
	// ============================================
	fmt.Println("=== POINTER TO STRUCT ===")
	cfg := config{host: "localhost", port: 8080, verbose: false}
	fmt.Printf("Before: %+v\n", cfg)

	updateConfig(&cfg, "0.0.0.0", 3000)
	fmt.Printf("After:  %+v\n", cfg)
	fmt.Println()

	// ============================================
	// ADDITION 6: Anonymous Struct
	// ============================================
	fmt.Println("=== ANONYMOUS STRUCT ===")
	temp := struct {
		x int
		y int
	}{x: 10, y: 20}
	fmt.Printf("Anonymous struct: %+v\n", temp)
	fmt.Println()

	// ============================================
	// ADDITION 7: Slice of Structs
	// ============================================
	fmt.Println("=== SLICE OF STRUCTS ===")
	students := []user{
		{name: "Sabuj", email: "sabuj@gmail.com"},
		{name: "Rahim", email: "rahim@gmail.com"},
		{name: "Karim", email: "karim@gmail.com"},
	}

	for i, s := range students {
		fmt.Printf("  %d. %s (%s)\n", i+1, s.name, s.email)
	}
	fmt.Println()

	// ============================================
	// ADDITION 8: Map of Structs
	// ============================================
	fmt.Println("=== MAP OF STRUCTS ===")
	products := map[string]product{
		"laptop": {name: "Laptop", price: 999.99, stock: 10},
		"mouse":  {name: "Mouse", price: 29.99, stock: 50},
		"keyboard": {name: "Keyboard", price: 79.99, stock: 30},
	}

	for key, p := range products {
		fmt.Printf("  %s -> ", key)
		p.display()
	}
	fmt.Println()

	// ============================================
	// ADDITION 9: Struct Comparison
	// ============================================
	fmt.Println("=== STRUCT COMPARISON ===")
	u1 := user{name: "Sabuj", email: "sabuj@gmail.com"}
	u2 := user{name: "Sabuj", email: "sabuj@gmail.com"}
	u3 := user{name: "Rahim", email: "rahim@gmail.com"}

	fmt.Println("u1 == u2:", u1 == u2) // true (same values)
	fmt.Println("u1 == u3:", u1 == u3) // false
	fmt.Println()

	// ============================================
	// ADDITION 10: Zero Value of Struct
	// ============================================
	fmt.Println("=== ZERO VALUE OF STRUCT ===")
	var empty user
	fmt.Printf("Empty struct: %+v\n", empty)
	fmt.Printf("Name: '%s' (empty string)\n", empty.name)
	fmt.Printf("Email: '%s' (empty string)\n", empty.email)
	fmt.Println()

	// ============================================
	// ADDITION 11: Updating Struct Fields
	// ============================================
	fmt.Println("=== UPDATING STRUCT FIELDS ===")
	person := user{name: "Sabuj", email: "old@gmail.com"}
	fmt.Printf("Before: %+v\n", person)

	person.email = "new@gmail.com" // update field
	person.name = "Sabuj Khan"     // update field
	fmt.Printf("After:  %+v\n", person)
	fmt.Println()

	// ============================================
	// ADDITION 12: Struct in Function (Value vs Pointer)
	// ============================================
	fmt.Println("=== STRUCT IN FUNCTION (VALUE vs POINTER) ===")
	p := product{name: "Phone", price: 599.99, stock: 25}

	// Pass by value - original unchanged
	modifyByValue(p)
	fmt.Println("After modifyByValue:", p.price)

	// Pass by pointer - original modified
	modifyByPointer(&p)
	fmt.Println("After modifyByPointer:", p.price)
}

// ============================================
// ADDITION 12: Helper Functions
// ============================================

func modifyByValue(p product) {
	p.price = 100.00 // Only modifies copy
	fmt.Println("  Inside modifyByValue:", p.price)
}

func modifyByPointer(p *product) {
	p.price = 100.00 // Modifies original
	fmt.Println("  Inside modifyByPointer:", p.price)
}
