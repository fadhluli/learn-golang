package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeAddressToIndonesia(address Address) {
	address.Country = "Indonesia"
}

func main() {

	address := Address{"Subang", "Jawa Barat", "Indonesia"}

	changeAddressToIndonesia(address)

	fmt.Println(address)
}
