package main

import "fmt"

func random() interface{} {
	return 10
}

func main() {
	// result := random()
	// resultString := result.(string)

	// fmt.Println(resultString)

	// resultInt := result.(int) // Panic
	// fmt.Println(resultInt)

	var random interface{} = random()
	switch value := random.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}
