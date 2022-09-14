package main

import "fmt"

func main() {
	// _ berfungsi untuk ignore sebuat pemanggilan dari parameter yang adas

	firstName, _, lastName := getFullName()
	fmt.Println(firstName, lastName)
}

func getFullName() (string, string, string) {
	return "Nur", "Kubbil", "Iman"
}
