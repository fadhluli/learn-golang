package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "fadhlul",
		"address": "irsyad",
	}

	fmt.Println(person)

	//jika ingin membuat "map" tanpa perlu menginisialisasikan valuenya terlebih dahulu
	book := make(map[string]string)

	book["title"] = "Harry Potter"
	book["author"] = "fadhlul irsyad"
	book["wrong"] = "ups"

	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)
}
