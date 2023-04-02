package main

import "fmt"

func biodata(firstNama string, middleName string, lastName string) {
	fmt.Println("Hay Kembali Lagi ", firstNama, middleName, lastName)
}

func main() {
	//namaDepan := "Muhammad" //Bisa juga Memakai variable untuk di masukan ke kelengkapanbiodata
	biodata("Muhammad", "Aulia", "Rahman")
}

// flow kode function untuk blok dalam pemanggilan main , paramater ruang untuk data isiandata dan disikan di func yang di buat biasanya (paramater tipedata)
// lalu di panggil di proses fmt println untuk paramater
// func main berfungsi untuk pemanggilan function
// parameter bisa di tambah kapan saja misal variable namaDepan di func main
