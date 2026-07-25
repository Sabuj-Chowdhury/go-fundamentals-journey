package main

import "fmt"

func main() {
	// ============================================
	// SECTION 1: ARRAYS
	// ============================================

	fmt.Println("========== ARRAYS ==========")

	// 1. Array Declaration - fixed size, same type
	var numbers [5]int
	fmt.Println("Empty array:", numbers) // [0 0 0 0 0]

	// 2. Array Initialization - declare and assign values
	var fruits [3]string = [3]string{"Apple", "Banana", "Cherry"}
	fmt.Println("String array:", fruits)

	// 3. Short declaration with initialization
	colors := [4]string{"Red", "Green", "Blue", "Yellow"}
	fmt.Println("Colors:", colors)

	// 4. Auto-size with ... (compiler counts elements)
	scores := [...]int{90, 85, 78, 92, 88}
	fmt.Println("Scores:", scores)
	fmt.Println("Array length:", len(scores))

	// 5. Accessing array elements (index starts at 0)
	fmt.Println("First fruit:", fruits[0])   // Apple
	fmt.Println("Last color:", colors[3])     // Yellow

	// 6. Modifying array elements
	colors[2] = "Purple"
	fmt.Println("Modified colors:", colors)

	// 7. Iterating array with for loop
	fmt.Println("\nIterating fruits:")
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("  Index %d: %s\n", i, fruits[i])
	}

	// 8. Iterating with range (Go style)
	fmt.Println("\nIterating scores with range:")
	for index, value := range scores {
		fmt.Printf("  Score at index %d = %d\n", index, value)
	}

	// 9. Array with only index (ignoring value)
	fmt.Println("\nJust indices:")
	for i := range colors {
		fmt.Printf("  Index: %d\n", i)
	}

	// 10. Multidimensional array (2D array)
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("\n2D Matrix:")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("  matrix[%d][%d] = %d\n", i, j, matrix[i][j])
		}
	}

	// 11. Array comparison (arrays can be compared in Go)
	a1 := [3]int{1, 2, 3}
	a2 := [3]int{1, 2, 3}
	a3 := [3]int{1, 2, 4}
	fmt.Println("\na1 == a2:", a1 == a2) // true
	fmt.Println("a1 == a3:", a1 == a3)  // false

	// ============================================
	// SECTION 2: SLICES
	// ============================================

	fmt.Println("\n\n========== SLICES ==========")

	// 1. Slice declaration - dynamic size
	// A slice is a reference to an underlying array
	var names []string
	fmt.Println("Empty slice:", names)       // []
	fmt.Println("Slice length:", len(names)) // 0

	// 2. Slice initialization with values
	languages := []string{"Go", "Python", "JavaScript", "Rust"}
	fmt.Println("Languages:", languages)

	// 3. Creating slice with make()
	// make(type, length, capacity)
	temps := make([]float64, 5)    // length=5, capacity=5
	fmt.Println("Temps:", temps)

	temps2 := make([]float64, 3, 10) // length=3, capacity=10
	fmt.Println("Temps2:", temps2)
	fmt.Println("Length:", len(temps2), "Capacity:", cap(temps2))

	// 4. Append elements to slice
	var numbers2 []int
	fmt.Println("\nAppending to slice:")
	for i := 1; i <= 5; i++ {
		numbers2 = append(numbers2, i*10)
		fmt.Printf("  Appended %d -> %v (len=%d, cap=%d)\n", i*10, numbers2, len(numbers2), cap(numbers2))
	}

	// 5. Append multiple elements at once
	extra := []int{50, 60, 70}
	numbers2 = append(numbers2, extra...)
	fmt.Println("After appending multiple:", numbers2)

	// 6. Slice from array
	arr := [5]int{10, 20, 30, 40, 50}
	sliceFromArr := arr[1:4] // index 1 to 3 (exclusive)
	fmt.Println("\nOriginal array:", arr)
	fmt.Println("Slice arr[1:4]:", sliceFromArr)

	// 7. Slice syntax variations
	fullSlice := arr[:]      // all elements
	firstTwo := arr[:2]     // index 0 to 1
	lastThree := arr[2:]    // index 2 to end
	fmt.Println("Full slice [:]:", fullSlice)
	fmt.Println("First two [:2]:", firstTwo)
	fmt.Println("Last three [2:]:", lastThree)

	// 8. Slice operations
	fruits2 := []string{"Mango", "Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("\nOriginal:", fruits2)
	fmt.Println("First 3:", fruits2[:3])
	fmt.Println("Middle:", fruits2[1:4])
	fmt.Println("Last 2:", fruits2[len(fruits2)-2:])

	// 9. Copy slice
	original := []int{1, 2, 3, 4, 5}
	copied := make([]int, len(original))
	copy(copied, original)
	fmt.Println("\nOriginal:", original)
	fmt.Println("Copied:", copied)

	// Modify copy - original stays unchanged
	copied[0] = 999
	fmt.Println("After modifying copy:")
	fmt.Println("  Original:", original)
	fmt.Println("  Copied:", copied)

	// 10. Delete element from slice (using append)
	deleteSlice := []int{10, 20, 30, 40, 50}
	fmt.Println("\nBefore delete:", deleteSlice)
	// Delete element at index 2 (value 30)
	deleteSlice = append(deleteSlice[:2], deleteSlice[3:]...)
	fmt.Println("After deleting index 2:", deleteSlice)

	// 11. Iterating slice
	fmt.Println("\nIterating languages:")
	for i, lang := range languages {
		fmt.Printf("  %d. %s\n", i+1, lang)
	}

	// 12. Slice with make and capacity
	fmt.Println("\nSlice capacity demo:")
	s := make([]int, 0, 5)
	fmt.Printf("  Empty: len=%d, cap=%d\n", len(s), cap(s))
	for i := 0; i < 8; i++ {
		s = append(s, i)
		fmt.Printf("  Append %d: len=%d, cap=%d\n", i, len(s), cap(s))
	}

	// 13. 2D Slice
	fmt.Println("\n2D Slice:")
	twoD := make([][]int, 3)
	for i := range twoD {
		twoD[i] = make([]int, 4)
		for j := range twoD[i] {
			twoD[i][j] = i*4 + j
		}
	}
	for i := range twoD {
		fmt.Printf("  %v\n", twoD[i])
	}

	// ============================================
	// SECTION 3: POINTERS
	// ============================================

	fmt.Println("\n\n========== POINTERS ==========")

	// 1. Pointer basics - a pointer stores memory address
	x := 10
	var ptr *int = &x // & gets the address of x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x (&x):", ptr)
	fmt.Println("Value at address (*ptr):", *ptr) // * dereferences (gets value)

	// 2. Modify value through pointer
	*ptr = 20
	fmt.Println("\nAfter *ptr = 20:")
	fmt.Println("  x is now:", x) // x changes to 20

	// 3. Pointer with short declaration
	y := 3.14
	py := &y
	fmt.Println("\ny:", y)
	fmt.Println("&y:", py)
	fmt.Println("*py:", *py)

	*py = 6.28
	fmt.Println("After *py = 6.28, y:", y)

	// 4. Pointer to array
	arr2 := [3]int{100, 200, 300}
	ptrArr := &arr2
	fmt.Println("\nArray:", arr2)
	fmt.Println("Pointer to array:", ptrArr)
	fmt.Println("Dereferenced:", *ptrArr)

	// Modify array through pointer
	(*ptrArr)[1] = 999
	fmt.Println("After modifying through pointer:", arr2)

	// 5. Pointer to slice
	slice := []int{1, 2, 3}
	ptrSlice := &slice
	fmt.Println("\nSlice:", slice)
	fmt.Println("Pointer to slice:", ptrSlice)

	*ptrSlice = append(*ptrSlice, 4, 5)
	fmt.Println("After append through pointer:", slice)

	// 6. new() function - allocates memory, returns pointer
	p := new(int)
	fmt.Println("\nnew(int) value:", *p)    // 0 (zero value)
	fmt.Println("new(int) address:", p)
	*p = 42
	fmt.Println("After *p = 42:", *p)

	// 7. Pointer in function (pass by reference)
	num := 100
	fmt.Println("\nBefore double:", num)
	double(&num)
	fmt.Println("After double:", num)

	// 8. Swap function using pointers
	a, b := 5, 10
	fmt.Println("\nBefore swap: a =", a, "b =", b)
	swap(&a, &b)
	fmt.Println("After swap: a =", a, "b =", b)

	// 9. Pointer to slice in function
	mySlice := []int{10, 20, 30}
	fmt.Println("\nBefore addElements:", mySlice)
	addElements(&mySlice, 40, 50)
	fmt.Println("After addElements:", mySlice)

	// 10. Nil pointer
	var nilPtr *int
	fmt.Println("\nNil pointer:", nilPtr) // <nil>
	// fmt.Println(*nilPtr) // This would panic!

	if nilPtr != nil {
		fmt.Println("Not nil")
	} else {
		fmt.Println("Pointer is nil!")
	}

	// 11. Pointer to struct (bonus - common use case)
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Sabuj", Age: 25}
	ptrPerson := &person
	fmt.Println("\nStruct via pointer:", ptrPerson)
	fmt.Println("Name:", ptrPerson.Name) // Go auto-dereferences
	fmt.Println("Age:", ptrPerson.Age)
	ptrPerson.Age = 26
	fmt.Println("Updated age:", person.Age)
}

// double takes a pointer and doubles the value
func double(n *int) {
	*n = *n * 2
}

// swap exchanges two values using pointers
func swap(a, b *int) {
	*a, *b = *b, *a
}

// addElements appends elements to a slice via pointer
func addElements(s *[]int, elements ...int) {
	*s = append(*s, elements...)
}
