package main

import "fmt"

func main() {

	person := map[string]string{
		"name": "affad",
	}

	books := [...]string{
		"harry potter",
		"jumanji",
		"game of thrones",
	}

	fmt.Println(person)
	fmt.Println(books)
}
