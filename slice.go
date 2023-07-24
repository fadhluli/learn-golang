package main

import "fmt"

func main() {

	var days = []string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	slice1 := days[1:4]

	fmt.Println("Ini arrayDays", days)
	fmt.Println("Ini slice", slice1)
	fmt.Println("Ini len slice", len(slice1))
	fmt.Println("Ini cap slice", cap(slice1))

	//klo append ke slice yg capacitynya masih ada maka value array aslinya juga ikut berubah
	//sebaliknya jika append ke slice capacitynya sudah penuh, maka akan membuat array baru sehingga
	//array aslinya tidak ikut berubah
	sliceAppend1 := append(slice1, "JumatBaru")

	fmt.Println("Ini sliceAppend1", sliceAppend1)
	fmt.Println("Ini array", days)

	//jika ingin membuat "slice" tanpa perlu menginisialisasikan valuenya terlebih dahulu
	makeSlice := make([]string, 0, 4)
	printSlice("b", makeSlice)
	makeSlice = days[1:3]
	makeSlice = append(makeSlice, "RabuBaru")

	fmt.Println("Ini makeSlice", makeSlice)
	fmt.Println("Ini arrayDays", days)

}

func printSlice(s string, x []string) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
