# Go Concurrency Practice

A beginner's guide to understanding **sequential vs concurrent** execution in Go (Golang).

---

## What You'll Learn

### 1. The Problem: Sequential Execution

When functions run **one after another**, the total time is the sum of all individual times.

```go
func main() {
    start := time.Now()

    transferFile()   // 3 seconds
    sendUrl()        // 1 second
    sendEmail()      // 1 second

    fmt.Println("Total time took to complete ", time.Since(start))
}
```

**Execution flow:**

```
transferFile()   ████░░░░░░░░░░░░░░░░  3s
sendUrl()        ░░░░░░░░░░░░░░░██░░░  1s
sendEmail()       ░░░░░░░░░░░░░░░░░██  1s
                                                  
Total: 5 seconds
```

Each function **blocks** until it finishes before the next one starts.

### 2. The Three Functions

```go
func transferFile() {
    fmt.Println("Transferring files .....")
    time.Sleep(3 * time.Second)  // simulates slow file transfer
    fmt.Println("transfer completed")
}

func sendUrl() {
    fmt.Println("Sending url ......")
    time.Sleep(1 * time.Second)
    fmt.Println("url sent!")
}

func sendEmail() {
    fmt.Println("Sending email ...")
    time.Sleep(1 * time.Second)
    fmt.Println("Email sent!")
}
```

Each function simulates work with `time.Sleep` - file transfer takes 3s, the others take 1s each.

### 3. The Solution: Goroutines

In Go, run functions concurrently with the `go` keyword - this creates a **goroutine**.

```go
go transferFile()
go sendUrl()
go sendEmail()
```

**With goroutines (concurrent):**

```
transferFile()   ████████████████████  3s (runs in background)
sendUrl()        ████░░░░░░░░░░░░░░░░  1s
sendEmail()      ████░░░░░░░░░░░░░░░░  1s

Total: 3 seconds  (instead of 5!)
```

> Note: The current code runs sequentially to demonstrate the problem. Using `go` keyword will be covered in the next lesson.

---

## Key Concurrency Facts

1. **Sequential** = total time is the sum of all function times
2. **Concurrent** = functions run in parallel, total time is the **longest** function
3. **Goroutines** are Go's lightweight threads managed by the Go runtime
4. **`go` keyword** starts a goroutine - the function runs in the background
5. **Channels** are used to communicate between goroutines (covered later)

---

## Expected Output

```
Transferring files .....
transfer completed
Sending url ......
url sent!
Sending email ...
Email sent!
Total time took to complete  5s
```

---

## How to Run

```bash
go run main.go
```

## Resources

- [Go by Example - Goroutines](https://gobyexample.com/goroutines)
- [Go by Example - Channels](https://gobyexample.com/channels)
- [Go Documentation](https://go.dev/doc/)

## Author

Sabuj - Learning Go!
