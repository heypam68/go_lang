package main

import "fmt"

func main() {
	nama := "Muhammad Aulia Rahman"
	perhitungan := 0

	penambahan := func() {
		nama := "Rahman"
		perhitungan++
		fmt.Println(nama)
	}

	penambahan()
	penambahan()
	penambahan()
	penambahan()

	fmt.Println(perhitungan)
	fmt.Println(nama)

	barang := "Laptop"
	kuantitasBarang := 0

	jumlahBarang := func() {
		barangLain := "Monitor"
		kuantitasBarang++
		fmt.Println(barangLain)
	}

	jumlahBarang()
	jumlahBarang()

	fmt.Println(kuantitasBarang)
	fmt.Println(barang)
}

//jadi closure adalah pembuatan data di dalam scope function ke function biasanya merubah variable yang sudah ada di atas
// harus membuat parammater untuk tidak merubah data baru "nama_variabel := "isian"
// Conting nama rahman di bawahnya penambahan
