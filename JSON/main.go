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

	p1 := Person{
		Name: "Sabuj",
		Age:  32,
		City: "Dhaka",
	}

	rawBytes, err := json.Marshal(p1)

	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println(string(rawBytes))

}
