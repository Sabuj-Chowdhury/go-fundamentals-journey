# Go Enums Practice

A beginner's guide to **enums** (enumerations) in Go (Golang).

---

## What You'll Learn

### 1. What is an Enum?

An enum represents a **fixed set of related constants**. Go doesn't have a dedicated `enum` keyword - instead, enums are created using **custom types** and **`iota`**.

```go
type WeekDays int

const (
    Monday    WeekDays = iota   // 0
    Tuesday                     // 1
    wednesday                   // 2
    Thursday                    // 3
    Friday                      // 4
    Saturday                    // 5
    Sunday                      // 6
)
```

`iota` is an auto-incrementing counter inside a `const` block.

### 2. Using `iota`

`iota` starts at `0` and increases by `1` for each constant in the block.

```go
const (
    A = iota   // 0
    B          // 1
    C          // 2
)
```

**How it works:**

```
iota block:
┌──────────┬───────┐
│ constant │ value │
├──────────┼───────┤
│ Monday   │   0   │
│ Tuesday  │   1   │
│ wednesday│   2   │
│ Thursday │   3   │
│ Friday   │   4   │
│ Saturday │   5   │
│ Sunday   │   6   │
└──────────┴───────┘
```

### 3. Using Enums with `switch`

```go
func getWorkingDay(day WeekDays) string {
    switch day {
    case Monday, Tuesday, wednesday, Thursday, Sunday:
        return "Office is open"
    case Friday, Saturday:
        return "Office is Closed"
    default:
        return "Invalid Day"
    }
}

result := getWorkingDay(Monday)
fmt.Println(result)  // Office is open
```

### 4. String-Based Enums

Enums can also use `string` instead of `int`.

```go
type status string

const (
    Open   status = "open"
    Closed status = "closed"
)
```

---

## Key Enum Facts

1. **`iota` resets to 0** at the start of each `const` block
2. **Enum values are just constants** - no special runtime support
3. **`iota` supports expressions:**
   ```go
   const (
       _  = iota             // skip 0
       KB = 1 << (10 * iota) // 1 << 10 = 1024
       MB                     // 1 << 20 = 1048576
       GB                     // 1 << 30
   )
   ```
4. **Unexported enum values** (like `wednesday`) are accessible within the same package but not outside

---

## Expected Output

```
Office is open
```

---

## How to Run

```bash
go run main.go
```

## Resources

- [Go by Example - Constants](https://gobyexample.com/constants)
- [Effective Go - iota](https://go.dev/doc/effective_go#iota)
- [Go Documentation](https://go.dev/doc/)

## Author

Sabuj - Learning Go!
