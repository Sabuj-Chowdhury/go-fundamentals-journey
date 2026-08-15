# Go Maps Practice

A beginner's guide to understanding Maps in Go (Golang).

---

## What You'll Learn

### 1. What is a Map?

A map is Go's **key-value** data structure (like a dictionary or hash table). Keys map to values for fast lookups.

```
Key   →  Value
"food"  →  "burger"
"drink" →  "coffee"
```

### 2. Creating a Map with `make()`

```go
// Creates an empty map
myMap := make(map[string]string)

// Adding key-value pairs
myMap["s1"] = "sabuj"
myMap["s2"] = "sabuj"

fmt.Println(myMap)  // map[s1:sabuj s2:sabuj]
```

**Structure:** `map[keyType]valueType`
- `map[string]string` → string keys, string values
- `map[string]int` → string keys, int values
- `map[int]string` → int keys, string values

### 3. Deleting from a Map

Use the built-in `delete()` function.

```go
delete(myMap, "s2")

fmt.Println(myMap)  // map[s1:sabuj]  (s2 removed)
```

**Syntax:** `delete(mapName, key)`

### 4. Creating a Map with Literal

```go
myMap2 := map[string]string{
    "first": "adip",
    "last":  "khadip",
}

fmt.Println(myMap2)
// map[first:adip last:khadip]
```

### 5. Map with Struct Values

Maps can store any type as values, including structs.

```go
type user struct {
    name  string
    email string
}

myMap3 := map[string]user{
    "data": user{
        name:  "Sabuj",
        email: "sabuj@gmail.com",
    },
}

// Access struct field through map key
fmt.Println(myMap3["data"].name)   // Sabuj
fmt.Println(myMap3["data"].email)  // sabuj@gmail.com
```

```
myMap3:
┌────────┬───────────────────────┐
│  key   │  value (struct)       │
├────────┼───────────────────────┤
│ "data" │  name:  "Sabuj"       │
│        │  email: "sabuj@gmail" │
└────────┴───────────────────────┘
```

---

## Key Operations Summary

| Operation | Syntax | Example |
|-----------|--------|---------|
| Create (empty) | `make(map[K]V)` | `make(map[string]int)` |
| Create (literal) | `map[K]V{...}` | `map[string]int{"a": 1}` |
| Add / Update | `m[key] = value` | `m["s1"] = "sabuj"` |
| Get | `m[key]` | `m["s1"]` → `"sabuj"` |
| Delete | `delete(m, key)` | `delete(m, "s2")` |
| Length | `len(m)` | `len(m)` → number of keys |

---

## Important Map Facts

1. **Maps are reference types** - no need for pointers, they're shared automatically
2. **Keys must be comparable** - use `string`, `int`, `bool`, etc. (not slices/maps)
3. **Default zero value is `nil`** - can't add to a nil map:
   ```go
   var m map[string]int
   // m["x"] = 1  ← PANIC! nil map
   // Must use make() first:
   m = make(map[string]int)
   m["x"] = 1  // OK
   ```
4. **Checking if a key exists:**
   ```go
   value, exists := myMap["s1"]
   fmt.Println(value, exists)  // sabuj true
   ```
5. **Iterating a map:**
   ```go
   for key, value := range myMap {
       fmt.Println(key, value)
   }
   ```

---

## Expected Output

```
map[s1:sabuj]map[first:adip last:khadip]Sabuj
```

> Note: Map iteration order is **not guaranteed** in Go, but with only one key the output is stable.

---

## How to Run

```bash
go run main.go
```

## Resources

- [Go by Example - Maps](https://gobyexample.com/maps)
- [Go Blog - Maps](https://go.dev/blog/maps)
- [Go Documentation](https://go.dev/doc/)

## Author

Sabuj - Learning Go!
