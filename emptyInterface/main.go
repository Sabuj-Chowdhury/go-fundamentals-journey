package main

import "fmt"

// func Print(data interface{}) {
// 	fmt.Println(data)

// }

// mordern go --> any = interface{}

// func Print(data any){

// }

// type assertion
// ok  idiom
func Process(data any) {
	strData, ok := data.(string) //asserting data as string

	if ok {
		fmt.Println(strData)

	}
	intData, ok := data.(int)
	if ok {
		fmt.Println(intData)
	}
}

func main() {

	// var data interface{}

	// data = "Sabuj"
	// fmt.Println(data)
	// data = 25
	// fmt.Println(data)

	// Print([]int{3, 5, 6})
	// Print("Sabuj Chowdhury")

	Process(0)

}
