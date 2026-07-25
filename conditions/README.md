# Go Conditions Practice

A beginner's practice repository for learning conditional statements in Go (Golang).

## What You'll Learn

### 1. if Statement
Executes a block of code only if a condition is true.

```go
if (age < 18) {
    fmt.Println("Too young")
}
```

### 2. else if Statement
Tests multiple conditions in sequence when the previous conditions are false.

```go
} else if (age >= 18 && age <= 30) {
    fmt.Println("Adult!")
}
```

### 3. else Statement
Executes when all previous conditions are false.

```go
} else {
    fmt.Println("Old")
}
```

### 4. Logical Operators
Go supports standard logical operators for combining conditions:

- `&&` - Logical AND (both conditions must be true)
- `||` - Logical OR (at least one condition must be true)
- `!` - Logical NOT (reverses the condition)

```go
age >= 18 && age <= 30  // age is between 18 and 30
```

### 5. Comparison Operators
- `==` - Equal to
- `!=` - Not equal to
- `<` - Less than
- `>` - Greater than
- `<=` - Less than or equal to
- `>=` - Greater than or equal to

## Code Structure

```go
age := 15

if (age < 18) {
    fmt.Println("Too young")         // Output: Too young
} else if (age >= 18 && age <= 30) {
    fmt.Println("Adult!")
} else if (age > 30 && age < 60) {
    println("Working")
} else {
    println("Old")
}
```

## How to Run

1. Make sure you have Go installed ([Download Go](https://go.dev/dl/))
2. Clone this repository
3. Run the program:

```bash
go run main.go
```

## Expected Output

```
Too young
```

## Resources

- [Go Documentation](https://go.dev/doc/)
- [Go by Example - If/Else](https://gobyexample.com/if-else)
- [Effective Go - Control Structures](https://go.dev/doc/effective_go#control-structures)

## Author

Sabuj - Learning Go!
