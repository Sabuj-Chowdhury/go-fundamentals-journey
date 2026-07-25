# Go Functions Practice

A beginner's practice repository for learning Go (Golang) functions.

## What You'll Learn

### 1. Named Functions
Standard functions in Go that can be called by their name.

```go
func main() {
    fmt.Println("Hello, World!")
}
```

### 2. IIFE (Immediately Invoked Function Expression)
Functions that execute as soon as they are defined, without needing to be called separately.

```go
func () {
    fmt.Println("IIFE function")
}()
```

### 3. IIFE with Parameters
IIFEs can also accept parameters, just like regular functions.

```go
func (name string) {
    fmt.Println("Hello ", name)
}("Sabuj")
```

### 4. Anonymous Functions
Functions without a name, stored in a variable and called later using that variable.

```go
result := func (num1 int, num2 int) int {
    return num1 + num2
}

fmt.Println(result(2, 5))
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
IIFE function
Hello  Sabuj
7
```

## Resources

- [Go Documentation](https://go.dev/doc/)
- [Go by Example - Functions](https://gobyexample.com/functions)
- [Effective Go - Functions](https://go.dev/doc/effective_go#functions)

## Author

Sabuj - Learning Go!
