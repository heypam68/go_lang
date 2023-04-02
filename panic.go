package main

import "fmt"

func endApp() {
	fmt.Println("aplikasi Selesai")
}

func selaluJalan() {
	fmt.Println("Function Selalu Dipanggil")
}

func testError(hang bool) {
	defer selaluJalan()
	if hang {
		panic("Function hang coba cari function lain")
	}
	fmt.Println("Function Kesatu di eksekusi")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("aplikasi error")
	}
	fmt.Println("aplikasi berjalan")
}

func main() {
	runApp(false)
	testError(true)
}
