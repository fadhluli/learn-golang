package main

import "fmt"

type Persond struct {
	Name, Gender string
}

func main() {
	person1 := Persond{"Fadhlul", "L"}
	//var person2 *Persond = &person1 //cara manual
	person2 := &person1

	fmt.Println(person1)
	fmt.Println(person2)

}
