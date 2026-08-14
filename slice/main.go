package main

import "fmt"

func main() {

	// defining array
	numbers := [6]int{7, 8, 9, 10, 11, 12}

	// making slice
	slice := numbers[1:4] // 1st index: excluding last index [:]-> all elements

	fmt.Println(slice)

	// adding elements in the slice
	slice = append(slice, 15)
	slice = append(slice, 16)
	slice = append(slice, 17)
	fmt.Println(slice)

	// slice has len() and cap()
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
