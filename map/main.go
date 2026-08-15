package main

import "fmt"

type user struct {
	name  string
	email string
}

func main() {
	// creating map with make()
	myMap := make(map[string]string)
	// adding value to map
	myMap["s1"] = "sabuj"
	myMap["s2"] = "sabuj"

	// deleting key from map
	delete(myMap, "s2")

	fmt.Printf("%+v", myMap)

	// creating map with litterels
	myMap2 := map[string]string{
		"first": "adip",
		"last":  "khadip",
	}
	fmt.Printf("%+v", myMap2)

	// map with struct
	myMap3 := map[string]user{
		"data": user{
			name:  "Sabuj",
			email: "sabuj@gmail.com",
		},
	}

	fmt.Println(myMap3["data"].name)

}
