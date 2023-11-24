package main

import "fmt"

type Man struct {
	Name string
}

// klo gk pakai asterisk (*) makan akan pass by value
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	affad := Man{"Affad"}
	affad.Married()
	fmt.Println(affad)
}
