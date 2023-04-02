package main

import "fmt"

func jenisBarang() (barangke1, barangke2, barangke3, barangke4 string) {
	barangke1 = "Laptop"
	barangke2 = "Kipas"
	barangke3 = "Komputer"
	barangke4 = "Speaker"

	return barangke1, barangke2, barangke3, barangke4
}

func jenisDagangan() (item1, item2, item3, item4, item5 string) {
	item1 = "Cilor"
	item2 = "Cilung"
	item3 = "Cipuk"
	item4 = "Cireng"
	item5 = "Cigong"

	return
}

func main() {
	barangke1, barangke2, barangke3, barangke4 := jenisBarang()
	fmt.Println(barangke1, barangke2, barangke3, barangke4)

	itema, itemb, itemc, itemd, iteme := jenisDagangan()
	fmt.Println(itema)
	fmt.Println(itemb)
	fmt.Println(itemc)
	fmt.Println(itemd)
	fmt.Println(iteme)

}

//Pembuatan jenisbarang ---> isian barang berbentuk string yang akan di kembalikan ke return
// Panggil jenis barang di func jenis barang di func main
// mencetak outputan barang dari jenis barang
// nama variabel di fungsi main untuk penyempanan bebas untuk itema , itemb, itemc dan seterusnya
