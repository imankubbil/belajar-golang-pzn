package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "kubbil" {
		return NotFoundError
	}
	return nil
}

func main() {
	err := GetById("")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not found")
		} else {
			fmt.Println("Unknown error")
		}
	}
}
