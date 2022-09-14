package main

import (
	"fmt"
)

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Diubah"
	fmt.Println(slice1)

	slice1[0] = "MeiMei"
	fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Disamber")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
}
