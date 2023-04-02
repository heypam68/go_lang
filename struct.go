package main

import "fmt"

type NamaBarang struct {
	Namakomponen   string
	Jeniskomponen  string
	Jumlahkomponen int
	Hargakomponen  int
}

type IdentitasPelamar struct {
	NamaLengkap    string
	JenisKelamin   string
	Alamat         string
	Agama          string
	JumlahKeluarga int
}

type Costumer struct {
	Name    string
	Address string
	Age     int
}

type AlatMasak struct {
	JenisAlatMasak    string
	JumlahAlatMasak   int
	KomponenAlatMasak string
	HargaAlatMasak    int
}

type BahanDapur struct {
	JenisMasakan    string
	HargaBahan      int
	KomponenMasakan string
	KebutuhanDasar  int
}

func main() {
	var hasilbarang NamaBarang

	hasilbarang.Namakomponen = "Apple"
	hasilbarang.Jeniskomponen = "Laptop"
	hasilbarang.Jumlahkomponen = 1
	hasilbarang.Hargakomponen = 5000000

	var CangkupanAlamat IdentitasPelamar

	CangkupanAlamat.NamaLengkap = "Muhammad Aulia Rahman"
	CangkupanAlamat.Alamat = "Jln Mekar Hurip"
	CangkupanAlamat.JenisKelamin = "Laki-laki"
	CangkupanAlamat.JumlahKeluarga = 6
	CangkupanAlamat.Agama = "Islam"

	var kumpulandata Costumer

	kumpulandata.Name = "Muhammad Aulia Rahman"
	kumpulandata.Address = "Jl.Mekar Hurip no 17"
	kumpulandata.Age = 28

	Bastian := Costumer{ //struct literals cara penulisan
		Name:    "Bastian",
		Address: "Jl.Mekar Hurip 17",
		Age:     34,
	}

	AkumulatifAlatMasak := AlatMasak{
		JenisAlatMasak:    "Mangkok",
		JumlahAlatMasak:   2,
		KomponenAlatMasak: "Besi",
		HargaAlatMasak:    200000,
	}

	Bahandapur := BahanDapur{
		JenisMasakan:    "Soto",
		HargaBahan:      20000,
		KomponenMasakan: "Santan",
		KebutuhanDasar:  15000,
	}

	fmt.Println(AkumulatifAlatMasak)
	fmt.Println(AkumulatifAlatMasak.JenisAlatMasak)
	fmt.Println(AkumulatifAlatMasak.JumlahAlatMasak)
	fmt.Println(AkumulatifAlatMasak.KomponenAlatMasak)
	fmt.Println(AkumulatifAlatMasak.HargaAlatMasak)

	fmt.Println(kumpulandata)

	fmt.Println(Bastian.Name)
	fmt.Println(Bastian.Address)
	fmt.Println(Bastian.Age)

	fmt.Println(hasilbarang)
	fmt.Println(hasilbarang.Namakomponen)
	fmt.Println(hasilbarang.Jeniskomponen)
	fmt.Println(hasilbarang.Jumlahkomponen)
	fmt.Println(hasilbarang.Hargakomponen)

	fmt.Println(CangkupanAlamat) //Struct mencangkup semua data

	fmt.Println(Bahandapur)
	fmt.Println(Bahandapur.HargaBahan)
	fmt.Println(Bahandapur.JenisMasakan)
	fmt.Println(Bahandapur.KebutuhanDasar)
	fmt.Println(Bahandapur.KomponenMasakan)

}

/**
struct digunakan untuk definisi barang tekait
dan nanti isinya di declarasikan pada func main
lalu print

hasil barang satu satu detail di outputnya
cangkupanalamat langsung outputan nya
*/
