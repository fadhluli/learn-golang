package main

import "fmt"

func main() {

	// counter := 1

	// for counter <=10 {
	// 	fmt.Println(counter)
	// 	counter++
	// }

	//for statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println(counter)
	}

	//for untuk pengulangan slice
	slice := []string{"affad", "fadhlul", "irsyad"}

	for sliceCount := 0; sliceCount < len(slice); sliceCount++ {
		fmt.Println(slice[sliceCount])
	}

	//For Range
	for i, value := range slice {
		fmt.Println(i, value)
	}

	//For Range tanpa index
	for _, value := range slice {
		fmt.Println(value)
	}

	//For Range pengulangan pada map
	person := make(map[string]string)
	person["name"] = "affad"
	person["age"] = "23"

	for i, value := range person {
		fmt.Println(i, value)
	}
}
