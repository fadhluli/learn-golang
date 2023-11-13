package main

import "fmt"

type HasName interface {
	getName()
}

type Person struct {
	name string
}

func (person Person) getName() {
	fmt.Println("Hello my name is fadhlul irsyadd")
}

func main() {
	fadhlul := Person{"affad"}
	fadhlul.getName()
}
