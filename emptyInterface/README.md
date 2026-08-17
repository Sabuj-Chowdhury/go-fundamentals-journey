# Go Empty Interface Practice

A beginner's guide to the **empty interface** (`interface{}` / `any`) in Go (Golang).

---

## What You'll Learn

### 1. What is the Empty Interface?

An empty interface has **no methods**, so **every type satisfies it**. It lets you work with values of unknown type.

```go
var data interface{}

data = "Sabuj"
fmt.Println(data)  // Sabuj

data = 25
fmt.Println(data)  // 25
```

In modern Go (1.18+), `any` is a shorthand for `interface{}`:

```go
// Both are the same:
func Print(data interface{}) {}
func Print(data any) {}
```

### 2. Type Assertion with the `ok` Idiom

The **ok idiom** safely checks the underlying type and avoids a panic if the assertion fails.

```go
func Process(data any) {
    strData, ok := data.(string)  // assert data as string
    if ok {
        fmt.Println(strData)
    }

    intData, ok := data.(int)     // assert data as int
    if ok {
        fmt.Println(intData)
    }
}
```

**How it works:**

```
data.(string)

    ok = false          ok = true
       │                    │
       ▼                    ▼
    skip               strData = the string
                         fmt.Println(strData)

Passing 0 (int):
    data.(string)  → ok = false  → skip
    data.(int)     → ok = true   → fmt.Println(0)
```

### 3. Why Use the `ok` Idiom?

A type assertion **without** `ok` will **panic** if the type doesn't match:

```go
// UNSAFE - panics if data is not a string
strData := data.(string)
```

```go
// SAFE - returns ok=false instead of panicking
strData, ok := data.(string)
if ok {
    fmt.Println(strData)
}
```

---

## When to Use the Empty Interface

| Scenario | Example |
|----------|---------|
| A function that accepts any type | `func Print(data any)` |
| Storing different types in a slice | `[]any{"hello", 42, true}` |
| JSON decoding (unknown structure) | `json.Unmarshal(data, &v)` |
| Generic libraries | Collections that hold any value |

---

## Expected Output

```
(no output)
```

> Passing `0` to `Process()` fails both the `string` and `int` assertions, so nothing is printed. Pass `42` or `"hello"` to see output.

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
