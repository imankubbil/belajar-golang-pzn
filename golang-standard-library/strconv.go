package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Result Bool:", result)
	}

	resultInt, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Result Int:", resultInt)
	}

	binary := strconv.FormatInt(898, 2)
	fmt.Println("binary:", binary)

	var stringInt string = strconv.Itoa(773)
	fmt.Println("stringInt:", stringInt)
}
