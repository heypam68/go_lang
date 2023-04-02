package main

import "fmt"

func Verbs() interface{} {
	// return 1
	// return true
	return "Verbs"
}

func Condition(a int) interface{} {
	if a == 1 {
		return 1
	} else if a == 10 {
		return true
	} else {
		return "Kondisi yang Anda Masukan Salah Coba Lagi"
	}
}

func PemilihanBarang(barang int) interface{} {
	if barang <= 20 {
		return false
	} else if barang >= 30 {
		return true
	} else {
		return "Stok Barang Kosong"
	}
}

func SortirBaju(baju int) interface{} {
	if baju == 0 {
		return "Baju kosong"
	} else if baju <= 20 {
		return "Stok ada"
	} else {
		return "Stok Barang Kosong"
	}
}

func PemilihanElektronik(elektronik bool) interface{} {
	if elektronik == true {
		return "Barang Ada Stok"
	} else {
		return "Barang Tidak Ada Stok"
	}
}

func main() {
	verbs := Verbs() //interface biasa
	fmt.Println(verbs)

	var condisi interface{} = Condition(2) //interface kosong
	fmt.Println(condisi)

	var hasilPemilihan interface{} = PemilihanBarang(10)
	fmt.Println(hasilPemilihan)

	var pemilihanBaju interface{} = SortirBaju(30)
	fmt.Println(pemilihanBaju)

	var elektronik interface{} = PemilihanElektronik(false)
	fmt.Println(elektronik)
}

/**
interface kosong harus di deklarsikan dengan interface{} = fungsi yang di bikin
*/
