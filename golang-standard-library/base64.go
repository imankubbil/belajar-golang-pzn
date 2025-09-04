package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Imankubbil"
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	decode, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decode))
	}
}
