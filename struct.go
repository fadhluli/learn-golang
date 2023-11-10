package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {

	var affad Customer
	affad.Age = 10
	affad.Name = "Fadhlul"
	affad.Address = "Jl. Melati Putih"

	affad.sayHello("David")
}
