# Go More Interface Practice

A beginner's guide to using Interfaces in Go with a real-world **payment system** example.

---

## What You'll Learn

### 1. Interfaces for Flexible Code

Interfaces let you write code that works with **many different implementations** of the same behaviour. Here, a single payment service can work with Bkash, Nagad, or a mock - just by switching the implementation.

```go
type PaymentMethod interface {
    pay(amount float64)
}
```

Any type that has a `pay(amount float64)` method satisfies the `PaymentMethod` interface.

### 2. Pointer Receivers in Interfaces

Note that the methods use **pointer receivers** (`*Bkash`, `*Nagad`). This means only pointers to these types satisfy the interface.

```go
type Bkash struct {
    apiKey string
}

func (bk *Bkash) pay(amount float64) {
    fmt.Printf("Paying %.2f tk with Bkash\n", amount)
}
```

### 3. Storing an Interface in a Struct

An interface can be a **struct field**, letting the struct depend on any implementation.

```go
type PaymentService struct {
    method PaymentMethod
}

func NewPaymentService(method PaymentMethod) *PaymentService {
    return &PaymentService{
        method: method,
    }
}

func (ps PaymentService) checkout() {
    ps.method.pay(10.0)
}
```

### 4. The Power of Mocking

Interfaces make **testing easy** - create a mock implementation instead of real payment logic.

```go
type MockPaymentMethod struct{}

func (mockPM MockPaymentMethod) pay(amount float64) {
    fmt.Println("testing payment method")
}
```

### 5. Putting It All Together

```go
func main() {
    // mockPm satisfies PaymentMethod, so it can be passed in
    mockPm := MockPaymentMethod{}

    paymentService := NewPaymentService(mockPm)
    paymentService.checkout()
}
```

**How it works:**

```
PaymentService
      │
      └── method: PaymentMethod (interface)
              │
              ├── *Bkash  → "Paying X tk with Bkash"
              ├── *Nagad  → "Paying X tk with Nagad"
              └── Mock    → "testing payment method"
```

To use Bkash or Nagad, just pass the correct implementation:

```go
bkash := &Bkash{apiKey: "sdfsfd"}
paymentService := NewPaymentService(bkash)
paymentService.checkout()

nagad := NewNagad("safsfsf")
paymentService := NewPaymentService(nagad)
paymentService.checkout()
```

---

## Key Interface Facts

1. **Constructor functions** (like `NewPaymentService`) return the struct and set the interface field
2. **Pointer receivers** mean the type must be passed as a pointer to satisfy the interface
3. **Mocking** lets you test logic without real side effects (like actual payments)
4. The calling code (`PaymentService`) never needs to know *which* payment method it's using

---

## Expected Output

```
testing payment method
```

---

## How to Run

```bash
go run main.go
```

## Resources

- [Go by Example - Interfaces](https://gobyexample.com/interfaces)
- [Effective Go - Interfaces](https://go.dev/doc/effective_go#interfaces)
- [Go Documentation](https://go.dev/doc/)

## Author

Sabuj - Learning Go!
