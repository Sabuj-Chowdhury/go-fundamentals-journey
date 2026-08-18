# Go Callback Functions Practice

A beginner's guide to **callback functions** in Go (Golang).

---

## What You'll Learn

### 1. What is a Callback Function?

A callback is a **function passed as an argument** to another function. It lets you decide *what* behaviour to use at call time.

```go
func calculate(a int, b int, operation func(x int, y int) int) int {
    return operation(a, b)
}
```

`operation` is the callback - any function that takes two ints and returns an int.

### 2. Creating Callbacks with Anonymous Functions

Define a callback inline using a **function literal**.

```go
add := func(a int, b int) int {
    return a + b
}

division := func(a int, b int) int {
    return a / b
}
```

### 3. Using the Callback

Pass the function as a regular argument.

```go
result := calculate(10, 20, add)
fmt.Println(result)       // 30

result2 := calculate(20, 10, division)
fmt.Println(result2)      // 2
```

**How it works:**

```
calculate(10, 20, add)
         │    │    │
         │    │    └── operation = add function
         │    └─────── b = 20
         └──────────── a = 10

         operation(10, 20) → 10 + 20 → 30


calculate(20, 10, division)
         │    │     │
         │    │     └── operation = division function
         │    └──────── b = 10
         └───────────── a = 20

         operation(20, 10) → 20 / 10 → 2
```

### 4. Callback with No Return Value

A callback can also have no return value at all.

```go
func process(sayHello func()) {
    sayHello()
}

greet := func() {
    fmt.Println("Hello")
}

process(greet)  // prints: Hello
```

---

## Key Callback Facts

1. **Type of the callback** is defined in the function signature (e.g. `func(int, int) int`)
2. **Anonymous functions** (function literals) are commonly used as callbacks
3. **Callbacks enable flexible, reusable code** - the same function can work with different behaviours
4. **Same idea as higher-order functions** in other languages (map, filter, reduce)

---

## Expected Output

```
30
2
```

---

## How to Run

```bash
go run main.go
```

## Resources

- [Go by Example - Functions](https://gobyexample.com/functions)
- [Go by Example - Closures](https://gobyexample.com/closures)
- [Go Documentation](https://go.dev/doc/)

## Author

Sabuj - Learning Go!
