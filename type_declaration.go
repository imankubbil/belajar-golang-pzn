package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var NoKtpAhmad NoKTP = "2121243121212134324"
	var marriedStatus Married = true
	fmt.Println(NoKtpAhmad)
	fmt.Println(marriedStatus)
}
