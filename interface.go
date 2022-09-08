package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

// Interface Kosong
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func main() {
	var iman Person
	iman.Name = "Nur Iman"

	cat := Animal{
		Name: "Kucing",
	}

	var kosong interface{} = Ups(2)

	sayHello(iman)
	sayHello(cat)
	fmt.Println(kosong)
}
