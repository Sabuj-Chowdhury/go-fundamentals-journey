# Go Structs Practice

A comprehensive guide to understanding Structs in Go (Golang).

---

## What You'll Learn

### 1. Basic Struct Definition & Initialization

A struct is a **custom data type** that groups related fields together.

```go
type user struct {
    name  string
    email string
}

// Named fields
sabuj := user{
    name:  "Sabuj",
    email: "sabuj@gmail.com",
}

// Print with all fields
fmt.Printf("%+v\n", sabuj)  // {name:Sabuj email:sabuj@gmail.com}

// Access field
fmt.Println(sabuj.email)     // sabuj@gmail.com
```

### 2. Nested Struct

A struct can contain another struct as a field.

```go
type address struct {
    city    string
    country string
}

type employee struct {
    name    string
    email   string
    workAdd address  // nested struct
}

emp := employee{
    name:    "Rahim",
    email:   "rahim@company.com",
    workAdd: address{city: "Dhaka", country: "Bangladesh"},
}

// Access nested fields
fmt.Println(emp.workAdd.city)     // Dhaka
fmt.Println(emp.workAdd.country)  // Bangladesh
```

### 3. Struct Embedding (Composition)

Go doesn't have inheritance, but supports **embedding** for code reuse. Embedded struct fields are **promoted** to the parent.

```go
type user struct {
    name  string
    email string
}

type manager struct {
    user   // embedded (no field name!)
    team   int
    branch string
}

mgr := manager{
    user:   user{name: "Karim", email: "karim@company.com"},
    team:   15,
    branch: "Dhaka",
}

// Access directly (promoted fields)
fmt.Println(mgr.name)    // Karim (promoted from user)
fmt.Println(mgr.email)   // karim@company.com (promoted)
fmt.Println(mgr.team)    // 15
```

### 4. Methods on Structs

Methods are functions that work on a specific type using **receivers**.

```go
type product struct {
    name  string
    price float64
    stock int
}

// Value receiver - can read but not modify
func (p product) display() {
    fmt.Printf("%s: $%.2f\n", p.name, p.price)
}

// Pointer receiver - can read AND modify
func (p *product) applyDiscount(percent float64) {
    p.price = p.price * (1 - percent/100)
}

p := product{name: "Laptop", price: 999.99}
p.display()            // Laptop: $999.99
p.applyDiscount(10)    // 10% off
p.display()            // Laptop: $899.99
```

**When to use pointer receiver?**
- When method modifies the struct
- When struct is large (avoid copying)
- When consistency is needed

### 5. Constructor Pattern

Go doesn't have constructors, but the convention is to use `NewXxx` functions.

```go
type account struct {
    owner   string
    balance float64
}

func NewAccount(owner string, initialBalance float64) *account {
    return &account{
        owner:   owner,
        balance: initialBalance,
    }
}

// Usage
acc := NewAccount("Sabuj", 1000.00)
```

### 6. Pointer to Struct

Use `&` to get a pointer. Access fields the same way (Go auto-dereferences).

```go
type config struct {
    host string
    port int
}

cfg := config{host: "localhost", port: 8080}
ptr := &cfg

// Both work the same way
fmt.Println(cfg.host)   // localhost
fmt.Println(ptr.host)   // localhost (auto-dereferenced)

// Modify via pointer
ptr.port = 3000
fmt.Println(cfg.port)   // 3000
```

### 7. Anonymous Struct

Create a struct without a type name (useful for temporary data).

```go
temp := struct {
    x int
    y int
}{x: 10, y: 20}

fmt.Printf("%+v\n", temp)  // {x:10 y:20}
```

### 8. Slice of Structs

```go
students := []user{
    {name: "Sabuj", email: "sabuj@gmail.com"},
    {name: "Rahim", email: "rahim@gmail.com"},
    {name: "Karim", email: "karim@gmail.com"},
}

for i, s := range students {
    fmt.Printf("%d. %s\n", i+1, s.name)
}
```

### 9. Map of Structs

```go
products := map[string]product{
    "laptop":   {name: "Laptop", price: 999.99, stock: 10},
    "mouse":    {name: "Mouse", price: 29.99, stock: 50},
    "keyboard": {name: "Keyboard", price: 79.99, stock: 30},
}

for key, p := range products {
    fmt.Printf("%s: $%.2f\n", key, p.price)
}
```

### 10. Struct Comparison

Structs with **same type** and **comparable fields** can be compared with `==`.

```go
u1 := user{name: "Sabuj", email: "sabuj@gmail.com"}
u2 := user{name: "Sabuj", email: "sabuj@gmail.com"}
u3 := user{name: "Rahim", email: "rahim@gmail.com"}

fmt.Println(u1 == u2)  // true (same values)
fmt.Println(u1 == u3)  // false
```

### 11. Zero Value of Struct

When no values are assigned, fields get their **zero values**.

```go
var empty user
fmt.Printf("%+v\n", empty)
// {name: email:}

// Empty string "" for string
// 0 for int
// 0.0 for float64
// false for bool
```

### 12. Updating Struct Fields

```go
person := user{name: "Sabuj", email: "old@gmail.com"}
person.email = "new@gmail.com"  // update
person.name = "Sabuj Khan"      // update
```

### 13. Pass by Value vs Pass by Pointer

```go
// By VALUE - function gets a COPY, original unchanged
func modifyByValue(p product) {
    p.price = 100.00  // Only modifies copy
}

// By POINTER - function gets ADDRESS, original modified
func modifyByPointer(p *product) {
    p.price = 100.00  // Modifies original
}

p := product{price: 599.99}
modifyByValue(p)
fmt.Println(p.price)  // 599.99 (unchanged)

modifyByPointer(&p)
fmt.Println(p.price)  // 100.00 (changed!)
```

---

## Quick Reference

| Concept | Syntax | Description |
|---------|--------|-------------|
| Define | `type T struct {}` | Create struct type |
| Create | `t := T{field: value}` | Initialize struct |
| Access | `t.field` | Get field value |
| Pointer | `t := &T{}` | Pointer to struct |
| Method | `func (t T) m()` | Value receiver |
| Method | `func (t *T) m()` | Pointer receiver |
| Embed | `type B struct { A }` | Embed struct A in B |
| Nested | `type B struct { a A }` | Named field of type A |

---

## When to Use Structs

- Group related data together
- Create custom types with behavior (methods)
- Model real-world entities (User, Product, Order)
- Pass complex data to functions
- Implement interfaces

## How to Run

```bash
go run main.go
```

## Resources

- [Go by Example - Structs](https://gobyexample.com/structs)
- [Go by Example - Methods](https://gobyexample.com/methods)
- [Effective Go - Structs](https://go.dev/doc/effective_go#structs)
- [Go Blog - Struct Embedding](https://go.dev/doc/effective_go#embedding)

## Author

Sabuj - Learning Go!
