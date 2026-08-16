# Go Range Practice

A beginner's guide to understanding the `range` keyword in Go (Golang).

---

## What You'll Learn

### 1. What is `range`?

The `range` keyword iterates over data structures like **slices, arrays, maps, and strings**. It returns the **index/key** and **value** for each iteration.

```
for index, value := range collection {
    // code here
}
```

### 2. Range over a Slice

```go
arr := []int{2, 3, 4, 5, 6}

for key, value := range arr {
    fmt.Println(key, value)
}
```

**How it works:**

```
arr := [2, 3, 4, 5, 6]

Index:   0    1    2    3    4
         ┌────┬────┬────┬────┬────┐
         │ 2  │ 3  │ 4  │ 5  │ 6  │
         └────┴────┴────┴────┴────┘

Iteration 1: key=0 value=2
Iteration 2: key=1 value=3
Iteration 3: key=2 value=4
Iteration 4: key=3 value=5
Iteration 5: key=4 value=6
```

### 3. Range over a Map

```go
myMap := map[string]string{
    "Name":    "Sabuj",
    "Success": "Ok",
}

for key, value := range myMap {
    fmt.Println(key, value)
}
```

- `key` → the map key
- `value` → the value associated with that key

> Note: Map iteration order is **not guaranteed** in Go.

### 4. Ignoring the Index / Key

If you only need the value, use `_` (blank identifier) to skip the index:

```go
for _, value := range arr {
    fmt.Println(value)
}
```

Similarly, if you only need the index:

```go
for key := range arr {
    fmt.Println(key)
}
```

---

## Range Syntax Summary

| Collection | Variables | Example |
|-----------|-----------|---------|
| Slice / Array | `index, value` | `for i, v := range arr` |
| Map | `key, value` | `for k, v := range myMap` |
| String | `index, rune` | `for i, r := range str` |

---

## Expected Output

```
0 2
1 3
2 4
3 5
4 6
```

---

## How to Run

```bash
go run main.go
```

## Resources

- [Go by Example - Range](https://gobyexample.com/range)
- [Effective Go - Range](https://go.dev/doc/effective_go#for)
- [Go Documentation](https://go.dev/doc/)

## Author

Sabuj - Learning Go!
