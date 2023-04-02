package main

import "fmt"

func main() {
	//Contoh untuk kondisi true
	name := "Muhammad Aulia Rahman"

	if name == "Muhammad Aulia Rahman" {
		fmt.Println("True")
	}

	//Contoh untuk kondisi false yang tidak di eksekusi kodenya karena 80 tidak sama dengan 100
	var penjualan int = 100000

	if penjualan == 80000 {
		fmt.Println("Tergantung Kondisi Parameter di atas")
	}

	//Else Expression string langsung ekseskusi benar
	penjualanHp := "Apple"

	if penjualanHp == "Apple" {
		fmt.Println("Benar")
	} else {
		fmt.Println("Salah")
	}

	//Else Expression int langsung di eksekusi salah
	var penjualanBuah int = 120000

	if penjualanBuah == 140000 {
		fmt.Println("Benar")
	} else {
		fmt.Println("Salah")
	}

	//Else if Percabangan ada 3 if --> else if ----> else

	var hargaJual int = 250000

	if hargaJual == 200000 {
		fmt.Println("Harga valid")
	} else if hargaJual == 180000 {
		fmt.Println("Mendekati Harga Asli")
	} else {
		fmt.Println("Harga Tidak Valid")
	}

	// Else if Percabangan untuk string 2 percabangan else if
	product := "Tidak Ada Barang"

	if product == "Laptop" {
		fmt.Println("Barang sama")
	} else if product == "Laptop Mac" {
		fmt.Println("Barang Serupa")
	} else if product == "Laptop China" {
		fmt.Println("Barang Grade A")
	} else {
		fmt.Println("Barang tidak sama")
	}

	// Short Length string else if
	var nama string = "Rendy Rilly Aprialia"

	if nama == "Rully" {
		fmt.Println("Rully Valid")
	} else if nama == "Budi" {
		fmt.Println("Budi Valid")
	} else if nama == "Amir" {
		fmt.Println("Ammir Valid")
	} else if nama == "Aulia" {
		fmt.Println("Aulia Valid")
	} else {
		fmt.Println("Nama Tidak ditemukan")
	}

	//Pembuatan length if else pengecekan short statment
	if length := len(nama); length < 7 {
		fmt.Println("Nama Terlalu Pendek")
	} else {
		fmt.Println("Nama Sudah Sesuai")
	}
}
