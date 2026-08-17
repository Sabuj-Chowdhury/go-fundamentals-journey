# Go Variadic Functions Practice

A beginner's guide to **variadic functions** in Go (Golang).

---

## What You'll Learn

### 1. What is a Variadic Function?

A variadic function accepts a **flexible number of arguments** of the same type. Inside the function, they're treated as a **slice**.

```go
func add(numbers ...int) int {
    sum := 0
    for _, value := range numbers {
        sum += value
    }
    return sum
}
```

The `...` before the type makes the parameter variadic.

### 2. Calling a Variadic Function

Pass any number of arguments - Go automatically packs them into a slice.

```go
fmt.Println(add(1, 2, 3, 4, 5))  // 15
```

**How it works:**

```
add(1, 2, 3, 4, 5)

numbers → []int{1, 2, 3, 4, 5}

Index:   0    1    2    3    4
         ┌────┬────┬────┬────┬────┐
         │ 1  │ 2  │ 3  │ 4  │ 5  │
         └────┴────┴────┴────┴────┘

1 + 2 + 3 + 4 + 5 = 15
```

### 3. Variadic Alongside Other Parameters

A variadic parameter must be the **last parameter** in the function signature.

```go
func greet(prefix string, students ...string) {
    for _, student := range students {
        fmt.Println(prefix, student)
    }
}

greet("welcome", "Sabuj", "Kabul", "Haem", "Malek", "Rohit")
```

**How it works:**

```
greet("welcome", "Sabuj", "Kabul", "Haem", "Malek", "Rohit")
         │
         └── prefix = "welcome"
                                  │
                                  └── students = ["Sabuj", "Kabul", "Haem", "Malek", "Rohit"]
```

---

## Key Variadic Function Facts

1. **Only one variadic parameter per function** - must be the last parameter
2. **Internally creates a slice** - `...int` behaves like `[]int` inside the function
3. **Can pass a slice directly** - use `...` to unpack a slice when calling:
   ```go
   nums := []int{1, 2, 3}
   fmt.Println(add(nums...))  // 6
   ```

---

## Expected Output

```
15
welcome Sabuj
welcome Kabul
welcome Haem
welcome Malek
welcome Rohit
```

---

## How to Run

```bash
go run main.go
```

## Resources

- [Go by Example - Variadic Functions](https://gobyexample.com/variadic-functions)
- [Effective Go - Variadic Functions](https://go.dev/doc/effective_go#variadic)
- [Go Documentation](https://go.dev/doc/)

## Author

Sabuj - Learning Go!
