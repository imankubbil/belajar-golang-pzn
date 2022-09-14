package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var iman Customer
	iman.Name = "Nur Iman"
	iman.Address = "Cilandak"
	iman.Age = 14

	fmt.Println(iman)

	joko := Customer{
		Name:    "Joko Anwar",
		Address: "Indonesia",
		Age:     30,
	}

	fmt.Println(joko)
}
