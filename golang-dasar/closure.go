package main

import "fmt"

func main() {

	name := "Nur"
	counter := 0

	increment := func() {
		name := "Lekha"

		fmt.Println(counter)
		fmt.Println(name)
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
