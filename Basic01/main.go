package main

import "fmt"


func main (){

	order := coffeeOrder("Sabuj","Latte",450)
	fmt.Println(order)


}

func coffeeOrder(name string, coffee string, price int) string{

	orderDetails :=fmt.Sprintf("Order for %s: %s coffee costs %d taka",name,coffee,price)
	return orderDetails
	
}