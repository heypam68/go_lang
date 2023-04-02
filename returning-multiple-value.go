package main

import "fmt"

func elektronik() (string, string, string, string) {
	return "Kipas", "AC", "Laptop", "Monitor"
}

func namaLengkap() (string, string, string) {
	return "Muhammad", "Aulia", "Rahman"
}

func main() {
	ke1, ke2, ke3, ke4 := elektronik()
	fmt.Println(ke1, ke2, ke3, ke4)

	firstname, middlename, lastname := namaLengkap()
	fmt.Println(firstname, middlename, lastname)

}

// jadi pengembalian dari return secara beberapa variable kita buat fun ---> pengembalian tipe data ---> return masukan tipe data
// func main buat ruang untuk data misalkan ke1 dan seterusnya lalu panggil --> func elektronik
// prosesnya print data untuk beberapa outputan
// misalkan jika middlename tidak di pakai atau ada paramater yang di tidak di pakai bisa di tambahkan _
