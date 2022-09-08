package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	iman := Man{"Iman"}
	iman.Married()

	fmt.Println(iman)
}
