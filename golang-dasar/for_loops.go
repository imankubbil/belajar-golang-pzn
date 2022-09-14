package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke : ", counter)
		counter++
	}

	// 2 statement
	for counter1 := 1; counter1 <= 10; counter1++ {
		fmt.Println("Perulangan1 ke : ", counter1)
	}

	// For Range
	names := []string{"Nur", "Iman", "Lekha", "Sholehati"}

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}
