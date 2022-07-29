package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Nur"
	middleName = "Kubbil"
	lastName = "Iman"

	return
}

func main() {
	a, b, c := getCompleteName()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
