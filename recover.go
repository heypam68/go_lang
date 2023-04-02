package main

import "fmt"

func hasilakhir() {
	fmt.Println("aplikasi defer terpanggil")
	pesan := recover()
	if pesan == nil {
		fmt.Println("Aplikasi Salah  :", pesan)
	}
	fmt.Println("Aplikasi Selesai")
}

func cobaRecover(coba bool) {
	defer hasilakhir()
	if coba {
		panic("aplikasi Tidak Jalan")
	}
	fmt.Println("aplikasi function Pertama di jalankan")
}

func hasilperbandingan() {
	fmt.Println("perbandingan ke defer dipanggil")
	perintah := recover()
	if perintah == nil {
		fmt.Println("perbandingan tidak sesuai :", perintah)
	}
	fmt.Println("perbandingan di eksekusi")
}

func perbandingan(banding bool) {
	defer hasilperbandingan()
	if banding {
		panic("perbandingan salah")
	}
	fmt.Println("perbandingan benar")
}

func aplikasiLain() {
	fmt.Println("Aplikasi lain telah di panggil")
	masukan := recover()
	if masukan == nil {
		fmt.Println("Aplikasi tidak sesuai : ", masukan)
	}
	fmt.Println("Aplikasi di Eksekusi")
}

func aplikasiPerbandingan(aplikasi bool) {
	defer aplikasiLain()
	if aplikasi {
		panic("Aplikasi Ada Kesahalan")
	}
	fmt.Println("Aplikasi Dipakai")
}

func nonKeuangan() {
	fmt.Println("Aplikasi non Keuangan telah di panggil")
	keuangan := recover()
	if keuangan == nil {
		fmt.Println("Aplikasi non Keuangan :", keuangan)
	}
	fmt.Println("Aplikasi non Keuangan Dipakai")
}

func aplikasiKeuangan(keuangan bool) {
	defer nonKeuangan()
	if keuangan {
		panic("Aplikasi Mengarah Ke Defer non Keuangan")
	}
	fmt.Println("Aplikasi Selesai Dipakai")
}

func main() {
	cobaRecover(true)
	fmt.Println("hello")

	perbandingan(true)
	aplikasiPerbandingan(false)

	aplikasiKeuangan(true)
}

// recover digunakan untuk menyimpanan massage dari panic dan di taro di function yang di panggil defernya
// penggunakan if di recover sangat di perlukan untuk penggunjian biasanya tanpa menggunakan recover kode tersebut tidak bisa dijalankan lagi
// dengan menggunakan recover maka kode pertama ---> defer ---> kondisi ---> eksekusi kode berikutnya

//recover penambahan pesan untuk panic error
