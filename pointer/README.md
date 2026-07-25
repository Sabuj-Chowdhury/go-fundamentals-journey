# Go Pointers Practice

A beginner's guide to understanding Pointers in Go (Golang).

---

## What You'll Learn

### 1. Basic Pointer Declaration

A pointer stores the **memory address** of a variable.

```go
a := 10
b := &a  // & gets the address of a

fmt.Println("a value:", a)      // 10
fmt.Println("b value:", b)      // 0xc0000103b8 (memory address)
```

### 2. Dereferencing

Use `*` to get the value stored at a pointer's address.

```go
b := &a
fmt.Println(*b)  // 10 (value at address)

// Modify value through pointer
*b = 1000
fmt.Println(a)   // 1000 (a is changed!)
```

### 3. Pass by Value (Without Pointer)

When you pass an array to a function **without** a pointer, Go creates a **copy**. Changes inside the function don't affect the original.

```go
func arrayModificationWithoutPointer(arr [5]int) {
    arr[0] = 1000  // Only modifies the copy
    fmt.Println("inside without pointer function:", arr)
}

arr := [5]int{1, 2, 3, 4, 45}
arrayModificationWithoutPointer(arr)
fmt.Println(arr)  // [1 2 3 4 45] - original unchanged!
```

### 4. Pass by Reference (With Pointer)

When you pass a **pointer to array** (`*[5]int`), the function modifies the **original** array.

```go
func arrayModificationWithPointer(arr *[5]int) {
    arr[0] = 1000  // Modifies original array
    fmt.Println("inside with pointer function:", arr)
}

arr := [5]int{1, 2, 3, 4, 45}
arrayModificationWithPointer(&arr)
fmt.Println(arr)  // [1000 2 3 4 45] - original changed!
```

### 5. Modify Variable via Pointer Function

A function can modify any variable by accepting a pointer to it.

```go
func changeValue(x *int) {
    *x = 1000  // Change value at address
}

y := 20
changeValue(&y)
fmt.Println(y)  // 1000
```

---

## Key Concepts

| Symbol | Name | Purpose |
|--------|------|---------|
| `&` | Address-of | Gets memory address of a variable |
| `*` | Dereference | Gets value stored at a pointer address |
| `*int` | Pointer type | Declares a pointer to int |
| `*[5]int` | Array pointer | Pointer to an array |

---

## Pass by Value vs Pass by Reference

```
Pass by Value (copy):          Pass by Reference (pointer):
┌─────────────┐                ┌─────────────┐
│  Original   │                │  Original   │
│  arr = [1,2]│                │  arr = [1,2]│
└──────┬──────┘                └──────┬──────┘
       │ copy                         │ address
       ▼                              ▼
┌─────────────┐                ┌─────────────┐
│   Copy      │                │  Pointer    │
│  [1,2]      │─────┐         │  *arr ──────┼──→ Original
└─────────────┘     │         └─────────────┘
                    │
                    ▼
              Changes lost!     Changes saved!
```

---

## Code Output

```
inside without pointer function: [1000 2 3 4 45]
[1 2 3 4 45]
a value  10
b value  0xc0000103b8
after change a value  100
after chnage a value b value is  100
after change a value  1000
after chnage a value b value is  1000
20
1000
```

---

## When to Use Pointers

- Modify original data in functions
- Avoid copying large structs/arrays
- Share data between functions efficiently
- Optional fields (nil = not set)

## How to Run

```bash
go run main.go
```

## Resources

- [Go by Example - Pointers](https://gobyexample.com/pointers)
- [Effective Go - Pointers](https://go.dev/doc/effective_go#pointers_to_interfaces)
- [Go Blog - Pointers](https://go.dev/doc/effective_go#pointers)

## Author

Sabuj - Learning Go!
