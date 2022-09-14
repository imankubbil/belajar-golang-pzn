package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Cilandak", "Jakarta", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Bandung"

	address2 = &Address{"Babelan", "Jawa Barat", "Indonesia"}

	*address3 = Address{"Malang", "Jawa Tengah", "Indonesia"}

	var address4 *Address = new(Address)
	address4.City = "Bekasi"

	fmt.Println(address1) // address 1 tidak berubah
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4) // mengembalikan pointer data kosong

	/*
		Tanda karakter & berfungsi untuk mengkopi sebuah (passing by value) yang dimana penyimpanan datanya tidak sama

		Tanda karakter * berfungsi untuk merubah sebuah data (passing by reference) yang dimana penyimpanan datanya sama
	*/
}
