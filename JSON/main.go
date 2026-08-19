package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {

	// p1 := Person{
	// 	Name: "Sabuj",
	// 	Age:  32,
	// 	City: "Dhaka",
	// }

	// rawBytes, err := json.Marshal(p1)

	// if err != nil {
	// 	fmt.Println("Error ", err)
	// }
	// fmt.Println(string(rawBytes))

	jsonText := `{"name":"Sabuj","age":32,"city":"Dhaka"}`

	var p2 Person

	err := json.Unmarshal([]byte(jsonText), &p2)

	if err != nil {
		fmt.Println("error : ", err)
	}

	fmt.Printf("%+v", p2)

}
