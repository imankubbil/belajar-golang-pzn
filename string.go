package main

import "fmt"

func main() {
	fmt.Println("Nur Iman")
	fmt.Println("Iman Kubbil")

	fmt.Println(len("Iman"))
	fmt.Println("Imankubbil"[0]) // hasilnya 73 karena hal itu byte code dari I
	fmt.Println("Imankubbil"[1]) // hasilnya 109 karena hal itu byte code dari m
}
