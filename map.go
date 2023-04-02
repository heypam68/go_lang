package main

import "fmt"

func main() {
	//Key string disebut maps misalnya mau nambah key langsung saja seperti contoh kewarganegaaran kalo di dalam maps pake : kalo mau menambahkan pake =
	biodata := map[string]string{
		"Nama":   "Muhammad Aulia Rahman",
		"Alamat": "Jln.Mekar Hurip no 17",
		"Agama":  "Islam",
	}

	biodata["Kewarganegaraan"] = "Indonesia"

	fmt.Println(biodata)
	fmt.Println(biodata["Nama"])
	fmt.Println(biodata["Alamat"])
	fmt.Println(biodata["Agama"])
	fmt.Println(biodata["Kewarganegaraan"])

	//Key Integer maps contoh keuangan gimana isian keynya bisa string int , bisa int string
	kasRt := map[string]int{
		"Modal Awal":  90000,
		"Modal Akhir": 120000,
	}

	kasRt["Biaya Pengeluaran"] = 180000

	fmt.Println(kasRt)
	fmt.Println(kasRt["Modal Awal"])
	fmt.Println(kasRt["Modal Akhir"])
	fmt.Println(kasRt["Biaya Pengeluaran"])

	//coba proses delete map dan add map
	buku := map[string]string{
		"Pengarang": "Jk Rowling",
		"Pembuat":   "Novel Hariri",
		"Judul":     "21 Century For World",
	}

	buku["Tittle"] = "21 Century" //proses add map
	fmt.Println(len(buku))        //Melihat panjang data

	delete(buku, "Tittle") //proses delete

	fmt.Println(buku)
	fmt.Println(buku["Pengarang"]) //proses print maps di isi key
	fmt.Println(buku["Pembuat"])
	fmt.Println(buku["Judul"])
	fmt.Println(buku["Tittle"])
	fmt.Println(len(buku)) //Melihat Panjang data lent

}
