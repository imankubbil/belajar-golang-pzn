package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Perulangan break ke", i)
	}

	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue
		}

		fmt.Println("Perulangan continue ke", j)
	}
}
