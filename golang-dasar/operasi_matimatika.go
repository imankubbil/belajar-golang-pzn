package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 20
	var c = a * b
	fmt.Println(c)

	//  Augmented Assignments
	a += 10
	b *= 10
	c /= 5
	fmt.Println(c)
}
