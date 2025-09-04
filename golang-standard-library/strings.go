package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("kubbil", "kub"))
	fmt.Println(strings.Split("kubbil", "kub"))
	fmt.Println(strings.ToLower("KUBBIL"))
	fmt.Println(strings.ToUpper("kubbil"))
	fmt.Println(strings.Trim("      Kubbil      ", " "))
	fmt.Println(strings.ReplaceAll("Nur iman kubbil", "kubbil", "Nur"))
}
