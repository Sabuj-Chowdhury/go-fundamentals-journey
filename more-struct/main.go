package main

import "fmt"

type Customer struct {
	name       string
	gender     string
	isLoggedIn bool
}

func main() {

	// construct new instance
	newCustomer := func(name string, gender string, isLoggedIn bool) Customer {

		return Customer{
			name:       name,
			gender:     gender,
			isLoggedIn: isLoggedIn,
		}

	}

	// making new customer using constructor function
	sabuj := newCustomer("Sabuj", "Male", false)

	fmt.Printf("%+v", sabuj)

	fmt.Println()
	sabuj.greetingCustomer()

	sabuj.login()

	fmt.Printf("%+v", sabuj)

}

// reciver function

func (c *Customer) login() {
	c.isLoggedIn = true
}

func (c Customer) greetingCustomer() {
	fmt.Println("Hello ", c.name)

}
