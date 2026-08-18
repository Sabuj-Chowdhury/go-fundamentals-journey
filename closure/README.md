# Go Closure Practice

A beginner's guide to **closures** in Go (Golang).

---

## What You'll Learn

### 1. What is a Closure?

A closure is a function that **captures and remembers variables** from its surrounding scope, even after the outer function has finished executing.

```go
func counter() func() int {
    count := 0

    return func() int {
        count++
        return count
    }
}
```

The inner function "closes over" `count` - it keeps a reference to it.

### 2. Using the Closure

```go
next := counter()

fmt.Println(next())  // 1
fmt.Println(next())  // 2
fmt.Println(next())  // 3
fmt.Println(next())  // 4
```

**How it works:**

```
counter() is called once:
┌─────────────────────────────────┐
│  count := 0                     │
│  return inner function ─────────┼──→  next()
└─────────────────────────────────┘

Each call to next():
  next()  →  count++  →  count=1  →  return 1
  next()  →  count++  →  count=2  →  return 2
  next()  →  count++  →  count=3  →  return 3
  next()  →  count++  →  count=4  →  return 4
```

`count` is not a local variable of `next()` - it **lives outside** `next()`, but `next()` remembers it.

### 3. Why Does This Work?

When `counter()` returns, its local variables normally die. But the returned inner function **holds a reference** to `count`, keeping it alive in memory.

```
After counter() returns:

  counter() frame    ── destroyed
  count = 0          ── still alive (referenced by next)

  next ──→ inner function ──→ count (shared, persistent)
```

Each time `next()` is called, it reads and updates the **same** `count`.

---

## Key Closure Facts

1. **A closure captures the variable, not the value** - changes to `count` are visible across calls
2. **Each call to the outer function creates a fresh closure** with its own independent copy:
   ```go
   a := counter()
   b := counter()
   a()  // 1
   a()  // 2
   b()  // 1  (separate count!)
   ```
3. **Closures are commonly used for** state management, counters, callbacks, and encapsulation

---

## Expected Output

```
1
2
3
4
```

---

## How to Run

```bash
go run main.go
```

## Resources

- [Go by Example - Closures](https://gobyexample.com/closures)
- [Effective Go - Closures](https://go.dev/doc/effective_go#functional-literals)
- [Go Documentation](https://go.dev/doc/)

## Author

Sabuj - Learning Go!
