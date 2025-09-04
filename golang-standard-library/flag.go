package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "host to connect to")
	var port *int = flag.Int("port", 8080, "port to connect to")
	var username *string = flag.String("username", "admin", "username to connect with")
	var password *string = flag.String("password", "admin", "password to connect with")
	flag.Parse()

	fmt.Println("host:", *host)
	fmt.Println("port:", *port)
	fmt.Println("username:", *username)
	fmt.Println("password:", *password)
}
