package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian tidak boleh nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	var contohError = errors.New("Ups Error")
	result, err := Pembagian(6, 2)

	if err == nil {
		fmt.Println("Hasil", result)
	} else {
		fmt.Println("Error", err.Error())
	}

	fmt.Println(contohError)
}
