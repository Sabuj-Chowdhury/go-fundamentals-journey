# Go Interface Practice

A beginner's guide to understanding Interfaces in Go (Golang).

---

## What You'll Learn

### 1. What is an Interface?

An interface is a **set of method signatures** that a type must implement. Any type that has those methods automatically satisfies the interface - no explicit `implements` keyword needed (implicit implementation).

```go
type Animal interface {
    speak()
}
```

**Structure:** `type Name interface { methods }`
- Any type with a `speak()` method is an `Animal`
- `Dog`, `Cat`, and `Human` all satisfy the `Animal` interface

### 2. Implementing an Interface

Define the methods on your types - that's it!

```go
type Dog struct{}
type Cat struct{}
type Human struct {
    name string
}

func (d Dog) speak() {
    fmt.Println("Woof woof!")
}

func (c Cat) speak() {
    fmt.Println("Meaw Meaw!")
}

func (h Human) speak() {
    fmt.Println("Hello there", h.name)
}
```

### 3. Using an Interface as a Parameter

An interface lets you pass **different types** to the same function.

```go
func makeSound(a Animal) {
    a.speak()
}

func main() {
    tommy := Dog{}
    pusyCat := Cat{}
    sabuj := Human{name: "Sabuj"}

    makeSound(tommy)    // Dog
    makeSound(pusyCat)  // Cat
    makeSound(sabuj)    // Human
}
```

**How it works:**

```
makeSound(Animal)
     │
     ├── Dog    → "Woof woof!"
     ├── Cat    → "Meaw Meaw!"
     └── Human  → "Hello there Sabuj"
```

The function only cares that the value can `speak()` - not what type it is.

---

## Key Interface Facts

1. **Interfaces are implemented implicitly** - no `implements` keyword
2. **Interfaces are satisfied structurally** - if a type has all the methods, it satisfies the interface
3. **A type can implement multiple interfaces**
4. **Interfaces can be embedded** in other interfaces to compose them

---

## Expected Output

```
Woof woof!
Meaw Meaw!
Hello there  Sabuj
```

---

## How to Run

```bash
go run main.go
```

## Resources

- [Go by Example - Interfaces](https://gobyexample.com/interfaces)
- [Effective Go - Interfaces](https://go.dev/doc/effective_go#interfaces)
- [Go Documentation](https://go.dev/doc/)

## Author

Sabuj - Learning Go!
