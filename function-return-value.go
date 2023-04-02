package main

import "fmt"

func getHello(name string) string {
	return "Hello" + name
}

func buahBuahan(namaBuahke1 string) string {
	return "Disini jual Buah" + namaBuahke1
}

func kumpulanElektronik(elektronik1 string) string {
	return "Barang ini" + elektronik1
}

func main() {
	hasil := getHello("Muhammad Aulia Rahman")
	fmt.Println(hasil)

	kumpulanBuah := buahBuahan("Kedondong")
	fmt.Println(kumpulanBuah)

	hasilelektronik := kumpulanElektronik("Apple")
	fmt.Println(hasilelektronik)
}

// membuat function dengan nama gethello terus isian parameter dengan nama: name terus string di luar untuk tipe data pengembalian return
// return apa yang mau di kembalikan jangan lupa panggil parameter
// function main panggil function gethello  , membuat ruang hasil dari proses function gethello lalu di tampilkan ke hasil
