# Go Structs - Advanced Practice

A beginner's guide to advanced struct patterns in Go: constructors and receiver methods.

---

## What You'll Learn

### 1. Struct Definition

```go
type Customer struct {
    name       string
    gender     string
    isLoggedIn bool
}
```

A struct groups related data fields. Here `Customer` has three fields: name, gender, and login status.

### 2. Constructor Function

Go has no built-in constructors, but the common pattern is a function that creates and returns a struct instance.

```go
newCustomer := func(name string, gender string, isLoggedIn bool) Customer {
    return Customer{
        name:       name,
        gender:     gender,
        isLoggedIn: isLoggedIn,
    }
}
```

**Why use a constructor?**
- Guarantees a valid/complete instance (all fields set)
- Single place to apply default values
- Easy to change how objects are created in one spot

> Note: This example uses an **anonymous function** stored in a variable. A more common convention is a named function like `func NewCustomer(...) Customer`.

### 3. Creating an Instance

```go
sabuj := newCustomer("Sabuj", "Male", false)

fmt.Printf("%+v\n", sabuj)
// {name:Sabuj gender:Male isLoggedIn:false}
```

`%+v` is a format verb that prints the **field names** along with values.

### 4. Receiver Methods

A **receiver** attaches a function to a struct type. Two types:

#### Value Receiver (read-only)

```go
func (c Customer) greetingCustomer() {
    fmt.Println("Hello", c.name)
}
```

- Works on a **copy** of the struct
- Can read fields but **cannot modify** the original

```go
sabuj.greetingCustomer()  // Hello Sabuj
```

#### Pointer Receiver (modifies struct)

```go
func (c *Customer) login() {
    c.isLoggedIn = true
}
```

- Works on the **original** struct (via memory address)
- **Can modify** fields permanently

```go
sabuj.login()

fmt.Printf("%+v\n", sabuj)
// {name:Sabuj gender:Male isLoggedIn:true}
```

---

## Value Receiver vs Pointer Receiver

```
Value Receiver (c Customer):      Pointer Receiver (c *Customer):
┌───────────────┐                 ┌───────────────┐
│  Copy         │                 │  Original     │
│  ┌─────────┐  │                 │  ┌─────────┐  │
│  │name     │  │  read only      │  │name     │  │
│  │gender   │  │                 │  │gender   │  │
│  └─────────┘  │                 │  │loggedIn │◄─┼── can modify!
│  changes LOST │                 │  └─────────┘  │
└───────────────┘                 └───────────────┘
```

| | Value Receiver | Pointer Receiver |
|---|---|---|
| Syntax | `func (c Customer) m()` | `func (c *Customer) m()` |
| Works on | Copy | Original |
| Can modify struct? | No | Yes |
| Use when | Read-only methods | Methods that change data |
| Memory | Copies whole struct | Passes address (cheaper for big structs) |

---

## Expected Output

```
{name:Sabuj gender:Male isLoggedIn:false}
Hello  Sabuj
{name:Sabuj gender:Male isLoggedIn:true}
```

---

## Common Patterns to Know

### Named Constructor (more conventional)

```go
type Customer struct {
    name string
    age  int
}

func NewCustomer(name string, age int) Customer {
    return Customer{name: name, age: age}
}
```

### When to use pointer receiver
- Method modifies the struct (like `login()`)
- Struct is large (avoid copying)
- You want consistency (don't mix both on the same type)

---

## How to Run

```bash
go run main.go
```

## Resources

- [Go by Example - Methods](https://gobyexample.com/methods)
- [Effective Go - Receivers](https://go.dev/doc/effective_go#receiver)
- [Go Documentation](https://go.dev/doc/)

## Author

Sabuj - Learning Go!
