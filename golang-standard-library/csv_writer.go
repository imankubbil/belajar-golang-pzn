package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Name", "Age", "Height", "Weight"})
	_ = writer.Write([]string{"John", "20", "1.80", "75"})
	_ = writer.Write([]string{"Mary", "18", "1.65", "70"})

	writer.Flush()
}
