package main

import "fmt"

func main() {

	name := "affad"

	switch name {
	case "affad":
		fmt.Println("you are smart")
	case "irsyad":
		fmt.Println("you are handsome")
	default:
		fmt.Println("Euwhh")
	}

	// swith short statement
	switch age := 5; age {
	case 5:
		fmt.Println("you are too young")
	case 20:
		fmt.Println("you are an adult")
	default:
		fmt.Println("you are a babby")
	}

	//switch tanpa kondisi
	switch grade := 100; {
	case grade >= 90:
		fmt.Println("you are clever")
	case grade <= 70:
		fmt.Println("you must learn")
	default:
		fmt.Println("You are idiot")
	}
}
