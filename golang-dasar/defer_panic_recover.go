package main

import "fmt"

func main() {

	// Defer function adalah function yang kita bisa jadwalkan untuk di eksekusi setelah sebuah function selesai di eksekusi
	// Defer function akan selalu di eksekusi walaupun terjadi error di function yang di eksekusi

	// Panic function adalah function yang bisa kita gunakan untuk menghentikan program
	// Panic function biasanya ketika terjadi error pada saat program kita berjalan
	// Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

	// Recover adalah function yang bisa kita gunakan untuk menangkap data panic
	// Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan

	// runApplication(0)
	runApp(false)
}

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result ", result)
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error")
	}

	fmt.Println("Aplikasi berjalan")
}
