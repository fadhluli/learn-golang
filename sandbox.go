package main

import "fmt"

func main() {
	var (
		firstName = "Fadhlul"
		lastName  = "Irsyad"
	)
	address := "Jakarta Barat"

	sayHello := "Hello nama saya " + firstName + " " + lastName + " saya tinggal di " + address

	fmt.Println(sayHello)

}
