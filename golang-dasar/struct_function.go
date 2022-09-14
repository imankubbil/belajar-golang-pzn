package main

import "fmt"

type CustomerPremium struct {
	Name, Address string
	Age           int
}

func (customer CustomerPremium) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var iman CustomerPremium
	iman.Name = "Nur Iman"
	iman.Address = "Cilandak"
	iman.Age = 14

	iman.sayHi("Nur Iman")
}
