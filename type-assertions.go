package main

import "fmt"

func random() any {
	return "affad"
}

func main() {
	result := random()
	resultString := result.(string)

	//switch val := result.(type) {
	//
	//}
	fmt.Println(resultString)
}
