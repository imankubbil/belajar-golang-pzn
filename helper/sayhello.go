package helper

import "fmt"

var version = "1.0.3" // variable ini tidak bisa di akses dari luar
var Application = "Belajar Golang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

// function ini tidak bisa di akses dari luar
func sayGoodbye(name string) {
	fmt.Println("Goodbyre", name)
}
