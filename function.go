package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		sayHello()
	}
	sayHelloTo("Fadhlul", "Irsyad")
	fmt.Println(getHello("Affad"))

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	//menghiraukan return value
	//firstName, _ := getFullName()
	//fmt.Println(firstName, _)

	fmt.Println(getFullName2())
	sumAll(1, 3, 4)

	//with slice as argument (variadic, varargs)
	numbers := []int{
		1, 4, 5,
	}
	sumAll(numbers...)

	//function as a value
	goodBye := getGoodBye
	goodBye("affad")

}

func sayHello() {
	fmt.Println("Hello World")
}

// function with parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// function return value
func getHello(name string) string {
	return "Hello " + name
}

// function return multiple value
func getFullName() (string, string) {
	return "Fadhlul", "Irsyad"
}

// function return named value
func getFullName2() (firstName, lastName string) {
	firstName = "fadhlul"
	lastName = "irsyad"

	return
}

// variadic, varargs
func sumAll(numbers ...int) {
	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println(total)
}

// function as a value
func getGoodBye(name string) {
	fmt.Println(name)
}
