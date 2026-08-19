# Go JSON Practice

A beginner's guide to **JSON marshaling** in Go (Golang).

---

## What You'll Learn

### 1. What is `json.Marshal`?

`json.Marshal` converts a Go struct into **JSON bytes**. Import `encoding/json` and call `json.Marshal()` on a struct.

```go
import "encoding/json"

p1 := Person{
    Name: "Sabuj",
    Age:  32,
    City: "Dhaka",
}

rawBytes, err := json.Marshal(p1)
```

### 2. Struct Tags

Struct tags define how a field maps to a JSON key. Without tags, the JSON key uses the **Go field name**. Tags let you customise it.

```go
type Person struct {
    Name string `json:"name"`    // JSON key → "name"
    Age  int    `json:"age"`     // JSON key → "age"
    City string `json:"city"`    // JSON key → "city"
}
```

### 3. Exported Fields Rule

`json.Marshal` only encodes **exported fields** (capital first letter). Lowercase fields are invisible to it.

```go
// WRONG - produces {} because fields are unexported
type Person struct {
    name string
    age  int
    city string
}

// CORRECT - uppercase field names
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    City string `json:"city"`
}
```

### 4. Converting Bytes to String

`json.Marshal` returns `[]byte`. Convert to string to print readable JSON.

```go
fmt.Println(string(rawBytes))
```

---

## Key JSON Facts

1. **Always check the error** - `json.Marshal` can fail if a value is not serializable
2. **Struct tags are optional** - without them, the field name is used as the JSON key
3. **Other useful tags:**
   | Tag | Effect |
   |-----|--------|
   | `json:"-"` | Skip this field entirely |
   | `json:"name,omitempty"` | Omit if zero value |
   | `json:"name,string"` | Marshal as a string |
4. **To unmarshal (JSON → struct):** use `json.Unmarshal(data, &myStruct)`

---

## Expected Output

```
{"name":"Sabuj","age":32,"city":"Dhaka"}
```

---

## How to Run

```bash
go run main.go
```

## Resources

- [Go by Example - JSON](https://gobyexample.com/json)
- [Go Blog - JSON](https://go.dev/blog/json)
- [Go Documentation](https://go.dev/doc/)

## Author

Sabuj - Learning Go!
