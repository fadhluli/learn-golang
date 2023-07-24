package main

import "fmt"

func main() {

	name := "affad"

	if name == "affad" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Hi, boleh kenalan?")
	}

	//if-short statement
	if lenName := len(name); lenName > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah pas")
	}
}
