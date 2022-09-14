package main

import "fmt"

func main() {
	name := "Iman"

	switch name {
	case "Iman":
		fmt.Println("Hello Iman")
	case "Kubbil":
		fmt.Println("Hello Iman")
	default:
		fmt.Println("Selamat Datang")
	}

	// Short Statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Telalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}
}
