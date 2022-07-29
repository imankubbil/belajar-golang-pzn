package main

import "fmt"

func main() {
	name := "Iman"

	if name == "Nur" {
		fmt.Println("Hello " + name)
	} else if name == "Nazar" {
		fmt.Println("Anda Kue Lebaran ?")
	} else {
		fmt.Println("Selamat Datang")
	}

	// Short Statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
