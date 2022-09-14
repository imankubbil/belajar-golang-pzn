package main

import "fmt"

func main() {
	const firstName string = "Nur"
	const lastName = "Iman"

	// error karena variable constant tidak dapat diubah
	// firstName = "Mahmud"
	// lastName = "Hasabih"

	// multiple constant
	const (
		myHandphone = "Samsung"
		myLaptop    = "Asus"
		value       = 50
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(myHandphone)
	fmt.Println(myLaptop)
	fmt.Println(value)
}
