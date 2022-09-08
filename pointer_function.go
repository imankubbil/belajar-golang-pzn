package main

import "fmt"

type Address1 struct {
	City, Province, Country string
}

func ChangeCountryToIndonesiaNotPointer(address Address1) {
	address.Country = "Indonesia"
}

func ChangeCountryToIndonesiaPointer(address *Address1) {
	address.Country = "Indonesia"
}

func main() {
	var alamat = Address1{
		City:     "Depok",
		Province: "Jawa Barat",
		Country:  "",
	}

	var alamat1 = Address1{
		City:     "Cikarang",
		Province: "Jawa Barat",
		Country:  "",
	}

	ChangeCountryToIndonesiaNotPointer(alamat)
	ChangeCountryToIndonesiaPointer(&alamat1)

	// tidak berubah karena pass by value
	fmt.Println(alamat)
	// tidak berubah karena pass by reference
	fmt.Println(alamat1)
}
