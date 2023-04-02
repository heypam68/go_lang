package main

import (
	"fmt"
)

func main() {
	//yang dicari dengan println slice1 length nya saja , kalo yang mau di cari kapasitasnya dengan len dan cap
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//Perubahan data array
	//months[5] = "aulia"
	//fmt.Println(slice1)

	//Perubahan data slice1 contoh
	//slice1[1] = "Juni Rubah Lagi"
	//fmt.Println(months)

	//Append digunakan untuk penambahan data sementara dari slice sedangkan array pertama dan slidenya tidak berubah
	var slice2 = months[4:7]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Muhammad")
	fmt.Println(slice3)
	//fmt.Println(months)
	//fmt.Println(slice2)

	//Perubahan data array di slice 3
	slice3[2] = "Muhammad Aulia Rahman"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	// Recomended menggunakan make newslice yang pertama 3 length itu harus sesuai dengan kebutuhan
	newSlice := make([]string, 3, 10)

	newSlice[0] = "Muhammad"
	newSlice[1] = "Aulia"
	newSlice[2] = "Rahman"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice dari new slice atas
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//Perbedaan Array dan slice
	iniArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	iniSlice := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
