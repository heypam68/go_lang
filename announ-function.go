package main

import "fmt"

type Blacklist func(string) bool
type FilterBuah func(string) bool //penyederhanaan menggunakan alias atau type untuk parameter
type FilterProduct func(string) bool

func blacklistUser(name string, blacklist Blacklist) {
	if blacklist(name) == true {
		fmt.Println("You are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func sortirBuah(buah string, filterbuah FilterBuah) {
	if filterbuah(buah) == true { //Pembuatan Kondisi
		fmt.Println("Ini Buah", buah)
	} else {
		fmt.Println("Ini Bukan Buah", buah)
	}
}

func sortirProduct(product string, filterproduct FilterProduct) {
	if filterproduct(product) == true {
		fmt.Println("Ini Product", product)
	} else {
		fmt.Println("Ini Bukan", product)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Admin"
	}

	filterbuah := func(buah string) bool { //kondisi case pertama untuk pengecekan buah
		return buah == "Mangga"
	}

	kumpulanProduct := func(product string) bool {
		return product == "Toshiba"
	}

	blacklistUser("Eko", blacklist) // outputan dari beberapa kondisi dan function
	blacklistUser("Admin", blacklist)

	sortirBuah("Mangga", filterbuah)
	sortirBuah("Apple", filterbuah)

	sortirProduct("Apple", kumpulanProduct)
	sortirProduct("Toshiba", kumpulanProduct)

	sortirBuah("Kendodong", func(buah string) bool { //Cara Pemanggilan Ke 2
		return buah == "Kedondong"
	})

	sortirBuah("Mangga", func(buah string) bool { //Cara Pemanggilan ke 2
		return buah == "Mangga"
	})

}
