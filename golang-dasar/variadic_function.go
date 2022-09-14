package main

import "fmt"

func main() {
	total := sumALL(10, 10, 20, 30, 40, 50)
	fmt.Println(total)

	numbers := []int{10, 50, 30, 70, 80, 30}
	totalNumbers := sumALL(numbers...)
	fmt.Println(totalNumbers)
}

func sumALL(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}
