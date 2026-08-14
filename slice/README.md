# Go Slices Practice

A beginner's guide to understanding Slices in Go (Golang).

---

## What You'll Learn

### 1. Creating a Slice from an Array

A slice is a **flexible, dynamic view** into an array. It doesn't own the data - it references a portion of the underlying array.

```go
// Original array
numbers := [6]int{7, 8, 9, 10, 11, 12}

// Create slice from index 1 to index 3 (last index is EXCLUDED)
slice := numbers[1:4]
fmt.Println(slice)  // [8 9 10]
```

**How slicing works:**

```
numbers := [6]int{7, 8, 9, 10, 11, 12}

Index:      0    1    2    3    4    5
           ┌────┬────┬────┬────┬────┬────┐
           │ 7  │ 8  │ 9  │ 10 │ 11 │ 12 │
           └────┴────┴────┴────┴────┴────┘
                    └─────────┘
              numbers[1:4] → [8 9 10]
                 ↑ start   ↑ end (exclusive)
```

**Slice syntax variations:**

| Expression | Result | Meaning |
|-----------|--------|---------|
| `numbers[1:4]` | `[8 9 10]` | Index 1 to 3 |
| `numbers[:]` | `[7 8 9 10 11 12]` | All elements |
| `numbers[:3]` | `[7 8 9]` | First 3 elements |
| `numbers[3:]` | `[10 11 12]` | From index 3 to end |

### 2. Appending Elements

Use the `append()` function to add elements to a slice. **Always reassign the result!**

```go
slice := numbers[1:4]       // [8 9 10]

slice = append(slice, 15)   // [8 9 10 15]
slice = append(slice, 16)   // [8 9 10 15 16]
slice = append(slice, 17)   // [8 9 10 15 16 17]
```

```
Before append:              After append:
┌─────────────┐            ┌──────────────────┐
│ [8 9 10]    │            │ [8 9 10 15 16 17]│
└─────────────┘            └──────────────────┘
```

### 3. Length (`len`) vs Capacity (`cap`)

```go
slice = append(slice, 15)
slice = append(slice, 16)
slice = append(slice, 17)

fmt.Println(len(slice))  // 6 (number of elements)
fmt.Println(cap(slice))  // 5 (available space in backing array)
```

```
Underlying array:  [7 8 9 10 11 12 ...]
                             │
Slice: [8 9 10 15 16 17]
        └──── len = 6 ────┘
        └──────── cap = 5+ ... ────┘
```

**Key difference:**
- `len()` - how many elements the slice currently holds
- `cap()` - how many elements the underlying array can hold before needing to grow

**Important:** When a slice grows beyond capacity, Go creates a **new, larger** backing array and copies elements over.

---

## Expected Output

```
[8 9 10]
[8 9 10 15 16 17]
6
6
```

> Note: The exact `cap()` value may vary depending on Go's growth strategy, but it will be >= `len()`.

---

## Slice vs Array - Quick Comparison

| Feature | Array | Slice |
|---------|-------|-------|
| Size | Fixed | Dynamic |
| Declaration | `[6]int` | `[]int` |
| Passed to functions | By value (copy) | By reference |
| Create from | Direct declaration | From array or `make()` |
| Most common in Go | Rarely | Everywhere! |

---

## Common Mistakes

1. **Forgetting to reassign append:**
   ```go
   // WRONG - slice not updated
   append(slice, 15)

   // CORRECT
   slice = append(slice, 15)
   ```

2. **Off-by-one in slicing:**
   ```go
   numbers[1:4]  // includes index 1, 2, 3 - NOT index 4!
   ```

3. **Slices share the backing array:**
   ```go
   a := []int{1, 2, 3, 4, 5}
   b := a[1:3]  // b sees [2 3]
   a[1] = 99    // b now sees [99 3] too!
   ```

---

## How to Run

```bash
go run main.go
```

## Resources

- [Go by Example - Slices](https://gobyexample.com/slices)
- [Go Blog - Slices: usage and internals](https://go.dev/blog/slices-intro)
- [Effective Go - Slices](https://go.dev/doc/effective_go#slices)

## Author

Sabuj - Learning Go!
