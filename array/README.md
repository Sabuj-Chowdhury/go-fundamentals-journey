# Go Arrays, Slices & Pointers Practice

A comprehensive guide to understanding Arrays, Slices, and Pointers in Go (Golang).

---

## 1. ARRAYS

Arrays in Go are **fixed-size** collections of elements of the **same type**. Once declared, the size cannot change.

### Declaration & Initialization

```go
// Declare an empty array of 5 integers (zero-valued)
var numbers [5]int
// Output: [0 0 0 0 0]

// Declare and initialize with values
var fruits [3]string = [3]string{"Apple", "Banana", "Cherry"}

// Short declaration
colors := [4]string{"Red", "Green", "Blue", "Yellow"}

// Auto-size with ... (compiler counts elements)
scores := [...]int{90, 85, 78, 92, 88}
```

### Key Points
- Array size is part of its type: `[3]int` and `[5]int` are **different types**
- Arrays are **value types** - copying creates a separate copy
- Index starts at **0**
- Default values are **zero values** (0 for int, "" for string, false for bool)

### Accessing & Modifying

```go
fruits[0]  // Access first element: "Apple"
fruits[2]  // Access third element: "Cherry"
colors[2] = "Purple"  // Modify element
len(scores)  // Get array length: 5
```

### Iterating Arrays

```go
// Traditional for loop
for i := 0; i < len(fruits); i++ {
    fmt.Println(fruits[i])
}

// Range-based (Go style) - gives index and value
for index, value := range scores {
    fmt.Printf("Index %d: %d\n", index, value)
}

// Just indices (ignore value with _)
for i := range colors {
    fmt.Println(i)
}
```

### Multidimensional Arrays

```go
matrix := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}

// Access: matrix[row][col]
matrix[0][1]  // 2
matrix[1][2]  // 6
```

### Array Comparison

```go
a1 := [3]int{1, 2, 3}
a2 := [3]int{1, 2, 3}
a1 == a2  // true (same size, same values)
```

---

## 2. SLICES

Slices are **dynamic**, flexible views into underlying arrays. They are the most commonly used collection type in Go.

### Slice vs Array
| Feature | Array | Slice |
|---------|-------|-------|
| Size | Fixed | Dynamic |
| Declaration | `[5]int` | `[]int` |
| Memory | Value type | Reference type |
| Length | Part of type | Can change |

### Creating Slices

```go
// From literal
languages := []string{"Go", "Python", "JavaScript"}

// Using make()
s := make([]int, 5)        // length=5, capacity=5
s := make([]int, 3, 10)    // length=3, capacity=10

// Empty slice
var names []string
```

### Length vs Capacity

```
Length:   Number of elements currently in the slice
Capacity: Number of elements the underlying array can hold

slice := make([]int, 3, 10)
         |-- length=3 --|
         |---------- capacity=10 ----------|
```

### Append Elements

```go
var s []int

// Single append
s = append(s, 1)        // [1]

// Multiple appends
s = append(s, 2, 3, 4)  // [1 2 3 4]

// Append another slice
extra := []int{5, 6}
s = append(s, extra...)  // [1 2 3 4 5 6]
```

**Important:** `append` may create a new underlying array if capacity is exceeded!

### Slice from Array

```go
arr := [5]int{10, 20, 30, 40, 50}

arr[1:4]   // [20 30 40]    - index 1 to 3
arr[:2]    // [10 20]       - first 2 elements
arr[2:]    // [30 40 50]    - from index 2 to end
arr[:]     // [10 20 30 40 50] - all elements
```

### Copy Slice

```go
original := []int{1, 2, 3}
copied := make([]int, len(original))
copy(copied, original)

// Or simpler:
copied := append([]int(nil), original...)
```

### Delete Element

```go
s := []int{10, 20, 30, 40, 50}

// Delete element at index 2 (value 30)
s = append(s[:2], s[3:]...)
// Result: [10 20 40 50]
```

### 2D Slice

```go
twoD := make([][]int, 3)
for i := range twoD {
    twoD[i] = make([]int, 4)
}
```

---

## 3. POINTERS

Pointers store **memory addresses** of variables. They allow direct memory manipulation and efficient data passing.

### Basic Pointer Concepts

```
Variable:  x = 10
Address:   &x = 0xc0000b2008
Pointer:   ptr = &x (stores the address)
Dereference: *ptr = 10 (gets the value at address)
```

### Declaration & Usage

```go
x := 10
ptr := &x       // ptr stores address of x
fmt.Println(ptr)  // 0xc0000b2008 (memory address)
fmt.Println(*ptr) // 10 (value at address)

*ptr = 20        // Modify value through pointer
fmt.Println(x)   // 20 (x is now 20)
```

### new() Function

```go
p := new(int)    // Allocates memory, returns pointer
fmt.Println(*p)  // 0 (zero value)
*p = 42
fmt.Println(*p)  // 42
```

### Pointer to Array

```go
arr := [3]int{100, 200, 300}
ptrArr := &arr

(*ptrArr)[1] = 999  // Modify through pointer
// arr is now [100 999 300]
```

### Pointer to Slice

```go
slice := []int{1, 2, 3}
ptrSlice := &slice

*ptrSlice = append(*ptrSlice, 4, 5)
// slice is now [1, 2, 3, 4, 5]
```

### Pointers in Functions

```go
// Without pointer (pass by value - original unchanged)
func double(n int) {
    n = n * 2  // Only modifies local copy
}

// With pointer (pass by reference - original modified)
func doublePtr(n *int) {
    *n = *n * 2  // Modifies original value
}

// Swap function
func swap(a, b *int) {
    *a, *b = *b, *a
}
```

### Nil Pointer

```go
var p *int
fmt.Println(p)  // <nil>

// Always check for nil before dereferencing!
if p != nil {
    fmt.Println(*p)
}
```

### Pointer to Struct

```go
type Person struct {
    Name string
    Age  int
}

person := Person{Name: "Sabuj", Age: 25}
ptr := &person

// Go auto-dereferences - no need for (*ptr).Name
fmt.Println(ptr.Name)  // "Sabuj"
ptr.Age = 26           // Directly modify
```

---

## Key Takeaways

| Concept | When to Use |
|---------|-------------|
| **Array** | Fixed-size collections, when size is known at compile time |
| **Slice** | Dynamic collections, most common choice in Go |
| **Pointer** | Modify variables in functions, avoid copying large data |

## Common Mistakes to Avoid

1. **Nil pointer dereference** - Always check `ptr != nil` before using `*ptr`
2. **Forgetting append returns** - `s = append(s, x)` (must reassign)
3. **Array size mismatch** - `[3]int` and `[4]int` are different types
4. **Slice sharing underlying array** - Modifying one slice may affect another

## How to Run

```bash
go run main.go
```

## Resources

- [Go Documentation](https://go.dev/doc/)
- [Go by Example - Arrays](https://gobyexample.com/arrays)
- [Go by Example - Slices](https://gobyexample.com/slices)
- [Go by Example - Pointers](https://gobyexample.com/pointers)
- [Effective Go - Arrays, Slices](https://go.dev/doc/effective_go#arrays)
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)

## Author

Sabuj - Learning Go!
