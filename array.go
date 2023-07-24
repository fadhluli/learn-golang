package main

import "fmt"

func main() {

	var name [2]string
	name[0] = "affad"
	name[1] = "irsyad"

	var student = [3]string{
		"imam",
		"andi",
		"fulham",
	}

	// [...] jika kita blom tahu kapasistas dari arraynya
	var hari = [...]string{
		"senin",
		"selasa",
		"rabu",
	}

	fmt.Println(name)
	fmt.Println(student)
	fmt.Println(hari)

	//len() untuk mengetahui panjang dari array bukan jumlah datanya
	fmt.Println(len(student))
}
