package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Nur"
	names[1] = "Iman"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values)
	fmt.Println(len(values)) // berapa panjang array
}
