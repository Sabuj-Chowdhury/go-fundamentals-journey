package main

import "fmt"

// func process(sayHello func()) {

// 	sayHello()
// }

func calculate(a int, b int, operation func(x int, y int) int) int {

	return operation(a, b)

}

func main() {

	// greet := func() {
	// 	fmt.Println("Hello")
	// }

	// process(greet)

	add := func(a int, b int) int {
		return a + b
	}

	division := func(a int, b int) int {
		return a / b

	}

	result := calculate(10, 20, add)
	fmt.Println(result)

	result2 := calculate(20, 10, division)
	fmt.Println(result2)

}
