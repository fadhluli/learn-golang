package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// normal function
//func blackList(name string) bool {
//	if name == "anjing" {
//		return true
//	} else {
//		return false
//	}
//}

func main() {

	blackList := func(name string) bool {
		return name == "anjing"
	}
	registerUser("affad", blackList)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
