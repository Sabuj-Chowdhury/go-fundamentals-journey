# Go Loops Practice

A beginner's practice repository for learning loops in Go (Golang).

## What You'll Learn

### 1. for Loop
Go only has one loop type - the `for` loop. It's similar to C-style loops with initialization, condition, and post statements.

```go
for i := 1; i <= 10; i++ {
    // code here
}
```

**Parts of a for loop:**
- `i := 1` - Initialization (runs once before loop starts)
- `i <= 10` - Condition (loop continues while this is true)
- `i++` - Post statement (runs after each iteration)

### 2. continue Statement
Skips the rest of the current iteration and jumps to the next one.

```go
if (i % 2 == 0) {
    continue  // Skip even numbers
}
```

### 3. break Statement
Exits the loop immediately when a condition is met.

```go
if (i == 9) {
    break  // Stop the loop when i equals 9
}
```

## Code Structure

```go
func makeCoffee(i int) {
    fmt.Println("Making coffee .....", i)
}

func main() {
    for i := 1; i <= 10; i++ {
        if (i % 2 == 0) {
            continue  // Skip even numbers
        }

        makeCoffee(i)  // Only called for odd numbers

        if (i == 9) {
            break  // Exit when i equals 9
        }
    }
}
```

## Flow Diagram

```
i=1 (odd) → makeCoffee(1) → continue
i=2 (even) → skip
i=3 (odd) → makeCoffee(3) → continue
i=4 (even) → skip
i=5 (odd) → makeCoffee(5) → continue
i=6 (even) → skip
i=7 (odd) → makeCoffee(7) → continue
i=8 (even) → skip
i=9 (odd) → makeCoffee(9) → break (exit loop)
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
Making coffee ..... 1
Making coffee ..... 3
Making coffee ..... 5
Making coffee ..... 7
Making coffee ..... 9
```

## Resources

- [Go Documentation](https://go.dev/doc/)
- [Go by Example - For](https://gobyexample.com/for)
- [Effective Go - For](https://go.dev/doc/effective_go#for)

## Author

Sabuj - Learning Go!
