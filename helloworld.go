package main

import "fmt"

func main() {
	fmt.Println("Satu =", 1)
	fmt.Println("Dua =", 2)
	fmt.Println("Tiga Koma Lima =", 3.5, 2, 3)
	fmt.Println("Benar =", true)
	fmt.Println("Salah =", false)
	fmt.Println("Nama saya adalah =", "Fadhlul Irsyad")
	fmt.Println("Panjang dari nama saya adalah", len("Fadhlul Irsyad"))

	// "String"[index] akan mengeluarkan dari byte dari huruf dari posisi index
	fmt.Println("Huruf ke 1 dari nama Fadhlul Irsyad adalah", "Fadhlul Irsyad"[0])
}
