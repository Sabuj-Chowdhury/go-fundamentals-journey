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

### 5. Anonymous Callback: Inline Function

Instead of assigning to a variable, pass the function **directly inline** to the caller.

```go
multiply := calculate(10, 20, func(x, y int) int {
    return x * y
})

fmt.Println(multiply)  // 200
```

**How it works:**

```
calculate(10, 20, func(x, y int) int { return x * y })
         │    │              │
         │    │              └── operation = inline anonymous function
         │    └───────────────── b = 20
         └────────────────────── a = 10

         operation(10, 20) → 10 * 20 → 200
```

This is the same as assigning the function to a variable first - just shorter when you only use it once.

### 6. Higher-Order Functions (HOF)

A **higher-order function** is any function that:
1. **Takes a function as an argument**, or
2. **Returns a function**

`calculate` is a higher-order function because it takes a callback. Go has many built-in HOFs:

```go
// map - transforms each element
nums := []int{1, 2, 3}
doubled := map nums, func(n int) int { return n * 2 }

// filter - keeps elements that match
evens := filter nums, func(n int) bool { return n%2 == 0 }
```

**Key difference:**

| Concept | Description |
|---------|-------------|
| Callback | A function passed as an argument (the *what*) |
| Higher-Order Function | A function that accepts/returns other functions (the *how*) |

A HOF is the **wrapper**, a callback is the **thing it wraps**.

```
HOF (Higher-Order Function)
  └── accepts a Callback
        └── uses it inside

calculate()  ← HOF
  └── operation()  ← Callback
```

In Go, all callback patterns use HOFs - they are two sides of the same coin.

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
200
```

---

## How to Run

```bash
go run main.go
```

## Resources

- [Go by Example - Functions](https://gobyexample.com/functions)
- [Go by Example - Closures](https://gobyexample.com/closures)
- [Effective Go - Functions](https://go.dev/doc/effective_go#functions)
- [Go Documentation](https://go.dev/doc/)

## Author

Sabuj - Learning Go!
