# Go Defer Practice

A beginner's guide to the `defer` keyword in Go (Golang).

---

## What You'll Learn

### 1. What is `defer`?

`defer` schedules a function to run **after the surrounding function returns**. The function call is deferred until the end, but its arguments are **evaluated immediately**.

```go
func exampleDefer() int {
    result := 10

    defer defferedFunction(result)  // captured: result = 10

    result += 10

    fmt.Println("In function :", result)  // prints: 20

    return result  // returns: 20
}
```

### 2. Execution Order

The deferred function runs **after** the return but **before** the function actually exits.

```go
func defferedFunction(a int) {
    fmt.Println("In deffered function: ", a)
}
```

**What happens step by step:**

```
1. result := 10

2. defer defferedFunction(result)
   └── argument is evaluated NOW → captured value = 10
   └── function call is scheduled for later

3. result += 10  →  result = 20

4. fmt.Println("In function :", result)  →  prints: In function : 20

5. return result  →  returns 20

6. deferred function runs  →  prints: In deffered function:  10

7. fmt.Println(20)  →  prints: 20
```

---

## Key Defer Facts

1. **Arguments are evaluated immediately** - `defer` captures the value, not a reference:
   ```go
   x := 10
   defer fmt.Println(x)  // prints 10, not whatever x is later
   x = 20
   ```
2. **Defers run in LIFO order** - last deferred, first executed:
   ```go
   defer fmt.Println("first")
   defer fmt.Println("second")
   defer fmt.Println("third")
   // prints: third, second, first
   ```
3. **Return values are set before defer runs** - a deferred function can modify named return values
4. **Common uses:** closing files, releasing locks, cleanup tasks

---

## Expected Output

```
In function : 20
In deffered function:  10
20
```

---

## How to Run

```bash
go run main.go
```

## Resources

- [Go by Example - Defers](https://gobyexample.com/defer)
- [Effective Go - Defer](https://go.dev/doc/effective_go#defer)
- [Go Documentation](https://go.dev/doc/)

## Author

Sabuj - Learning Go!
