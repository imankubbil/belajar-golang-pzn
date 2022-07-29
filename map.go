package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Nur Iman",
		"address": "Jakarta",
	}

	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book = make(map[string]string)
	book["title"] = "Belajar Go-lang"
	book["author"] = "Eko"
	book["ups"] = "salah"
	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}
