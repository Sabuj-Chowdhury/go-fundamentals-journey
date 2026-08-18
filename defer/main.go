package main

import "fmt"

func defferedFunction(a int) {
	fmt.Println("In deffered function: ", a)
}

func exampleDefer() int {
	result := 10

	defer defferedFunction(result)

	result += 10

	fmt.Println("In function :", result)

	return result
}

func main() {

	fmt.Println(exampleDefer())

}
