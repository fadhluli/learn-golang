package main

import "fmt"

func main() {

	//menjadikan "aString" sebagai alias dari "string"
	type aString string

	var name aString = "Affad"

	fmt.Println(name)
}
