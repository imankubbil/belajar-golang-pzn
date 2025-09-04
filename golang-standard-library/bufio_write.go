package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	_, _ = writer.WriteString("Hello World\n")
	_, _ = writer.WriteString("Selamat Belajar\n")
	writer.Flush()
}
