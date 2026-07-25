package main

import "fmt"

// named functions (standred function)
func main(){
	// IIFE(Imedieatly Invoked Function Expression)
	func ()  {
		fmt.Println("IIFE function")
	}()

	// IIFE with parametar
	func (name string)  {
		fmt.Println("Hello ",name)
		
	}("Sabuj")

	// annoynums function
	result := func (num1 int, num2 int) int {
		return num1 + num2		
	}

	fmt.Println(result(2,5))
}
