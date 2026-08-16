package main

import "fmt"

func main() {
	// myMap := map[string]string{
	// 	"Name":    "Sabuj",
	// 	"Success": "Ok",
	// }

	// for key, value := range myMap {
	// 	fmt.Println(key, value)
	// }

	// for array
	arr := []int{2, 3, 4, 5, 6}
	for key, value := range arr {
		fmt.Println(key, value)
	}
}
