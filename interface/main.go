package main

import "fmt"

type Animal interface {
	speak()
}

type Dog struct{}
type Cat struct{}
type Human struct {
	name string
}

func (d Dog) speak() {
	fmt.Println("Woof woof!")
}
func (c Cat) speak() {
	fmt.Println("Meaw Meaw!")
}
func (h Human) speak() {
	fmt.Println("Hello there ", h.name)
}

func makeSound(a Animal) {
	a.speak()
}

func main() {
	tommy := Dog{}
	pusyCat := Cat{}
	sabuj := Human{
		name: "Sabuj",
	}

	makeSound(tommy)
	makeSound(pusyCat)
	makeSound(sabuj)
}
