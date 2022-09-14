package main

import "fmt"

func main() {
	result := getHello("Iman")
	fmt.Println(result)
}

func getHello(name string) string {
	return "Hello " + name
}
