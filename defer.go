package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(nilai int) {
	// defer logging() //digunakan untuk pemanggilan kode secara bertahap misalkan fungsi sudah di eksekusi lalu defer ke function lainnya yang di panggil oleh defer
	fmt.Println("Run Application")
	hasil := 10 / nilai
	fmt.Println("hasil", hasil)
}

func main() {
	runApplication(1)
}

// defer akan memanggil eksekusi selanjutnya walaupun kode tersebut error
// defer digunakan untuk memanggil function tertentu walaupun terjadi error untuk function lainnya
