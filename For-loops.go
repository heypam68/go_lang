package main

import "fmt"

func main() {
	//For loop angka++ simple
	angka := 10

	for angka <= 10 {
		fmt.Println("Angka Benar", angka)
		angka++
	}

	//For loop buah penjabaran dari ++
	buah := 152

	for buah <= 200 {
		fmt.Println("Buah Ke", buah)
		buah = buah + 1
	}

	//Short Statement atau init statement dan post statement dalam for loop gk harus bikin variable atau ruang lagi
	//uang := nomimal ---> init statement , uang++ ---> post statement , uang<= ---> kondisi , output ---> nominal ke uang

	for uang := 1000000; uang <= 1000010; uang++ {
		fmt.Println("Nominal Uang Ke", uang)
	}

	//For loops untuk slice contoh untuk iterasi data outputan manual memakai println slice
	nama := []string{"Muhammad", "Aulia", "Rahman"}

	//fmt.Println(nama[0]) //cara Manual
	//fmt.Println(nama[1]) //cara Manual
	//fmt.Println(nama[2]) //cara Manual

	//kita coba pake iterasi dengan for loops i variable wadah terserah mau di pake apa aj
	for i := 0; i < len(nama); i++ {
		//fmt.Println(i)
		fmt.Println(nama[i])
	}

	//For Range Maps slice gabungan
	elektronik := []string{"Laptop", "Speaker", "Monitor", "Cpu"}

	for i := 0; i < len(elektronik); i++ {
		//fmt.Println(i)
		fmt.Println(elektronik[i])
	}

	for index, kumpulan := range elektronik {
		fmt.Println("Katalog", index, "=", kumpulan)
	}

	//Contoh agar index nya gk mau ke print dari for range elektronik jadi langsung ke barang jadi index di inialisasi kosong _ (tidak dibutuhkan) dan bikin ruang katalog untuk penampungan elektronik
	for _, katalog := range elektronik {
		fmt.Println(katalog)
	}

	//Contoh gabungan for range di maps for key value
	biodata := make(map[string]string)
	biodata["Nama"] = "Muhammad Aulia Rahman"
	biodata["Alamat"] = "Jalan Mekar Hurip No 17"
	biodata["Gelar"] = "D3"

	for key, value := range biodata {
		fmt.Println(key, "=", value)
	}

	//Contoh Gabungan di for range maps for key value
	//person := make(map[string]string)
	//person["Nama"] = "Muhammad Aulia Rahman"
	//person["Alamat"] = "Jalan Mekar Hurip no 17"
	//
	//for key, value := range person {
	//	fmt.Println(key, "=", value)
	//}
}
