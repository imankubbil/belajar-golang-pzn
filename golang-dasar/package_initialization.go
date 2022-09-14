package main

import (
	"fmt"
	"programmer-zaman-now/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
