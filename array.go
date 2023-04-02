package main

import "fmt"

func main() {
	//Basic Array Nama Tempat String
	var namaTempat [5]string
	namaTempat[0] = "Bandung"
	namaTempat[1] = "Jakarta"
	namaTempat[2] = "Jogya"
	namaTempat[3] = "Aceh"
	namaTempat[4] = "Surabaya"

	fmt.Println(namaTempat[0])
	fmt.Println(namaTempat[1])
	fmt.Println(namaTempat[2])
	fmt.Println(namaTempat[3])
	fmt.Println(namaTempat[4])

	// Array Basic Uang Int
	var kumpulanUang [10]int
	kumpulanUang[0] = 100000
	kumpulanUang[1] = 200000
	kumpulanUang[2] = 300000
	kumpulanUang[3] = 400000
	kumpulanUang[4] = 500000
	kumpulanUang[5] = 600000
	kumpulanUang[6] = 700000
	kumpulanUang[7] = 800000
	kumpulanUang[8] = 900000
	kumpulanUang[9] = 1000000

	fmt.Println(kumpulanUang[0])
	fmt.Println(kumpulanUang[1])
	fmt.Println(kumpulanUang[2])
	fmt.Println(kumpulanUang[3])
	fmt.Println(kumpulanUang[4])
	fmt.Println(kumpulanUang[5])
	fmt.Println(kumpulanUang[6])
	fmt.Println(kumpulanUang[7])
	fmt.Println(kumpulanUang[8])
	fmt.Println(kumpulanUang[9])

	//Deklarasi Langsung Array nilai int
	var nilai = [10]int{
		90,
		80,
		80,
		80,
		70,
		60,
		75,
		85,
		90,
		90,
	}
	fmt.Println(nilai)

	//Deklarasi Langsung string
	var kota = [4]string{
		"Bandung",
		"Jakarta",
		"Surabaya",
	}
	fmt.Println(kota)

	// Deklarasi Len untuk mengetahui jumlah data array
	var kumpulanPoint = [4]int{
		10,
		20,
		30,
		40,
	}
	fmt.Println(len(kumpulanPoint))

	// Contoh dari Programer jaman now [...] ---> ini untuk tipikal array tak terhingga values[.] = yang mau di gantikan
	var values = [...]int{
		90,
		85,
		95,
		90,
		100,
		120,
		130,
		140,
		150,
		160,
		170,
		180,
		190,
		1290,
		132,
		32,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	values[15] = 1
	fmt.Println(values)
}
