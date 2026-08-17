package main

import "fmt"

func add(numbers ...int) int {
	sum := 0

	for _, value := range numbers {

		sum += value
	}
	return sum
}

func greet(prefix string, students ...string) {
	for _, student := range students {
		fmt.Println(prefix, student)
	}

}

func main() {

	fmt.Println(add(1, 2, 3, 4, 5))

	greet("welcome", "Sabuj", "Kabul", "Haem", "Malek", "Rohit")

	// variadic args
	students := []string{"Sabuj", "Kabul", "Haem", "Malek", "Rohit"}
	greet("Hello", students...)

}

/*
variadic functions benifits
1. flexiable amoount of args
2. makes internally a slice
3. has to be the only parameter or last perameter
*/
