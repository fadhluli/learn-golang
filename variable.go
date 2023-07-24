package main

import "fmt"

func main() {
	var name string

	name = "affad"
	fmt.Println(name)
	name = "Irsyad"
	fmt.Println(name)

	//jika mau deklarasi variable tanpa "var"
	myCat := "Meoww"
	fmt.Println(myCat)

	//deklarasi multiple variable
	var (
		firstName = "Fadhlul"
		lastName  = "Irsyad"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
