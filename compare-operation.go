package main

import "fmt"

func main() {

	// Perbandingan integer bool
	var nilaiA = 80
	var nilaiB = 90

	var result bool = nilaiB < nilaiA

	fmt.Println(result)

	//Perbanding string bool
	var tempatA = "Bandung"
	var tempatB = "Surabaya"

	var pilihanTempat bool = tempatA == tempatB

	fmt.Println(pilihanTempat)

	//Langsung di output != nilai kalo ini biasanya kebalikan jika hasilnya true maka false, berbeda kalo == apakah sama

	var nilai1 = 1000
	var nilai2 = 1200

	fmt.Println(nilai1 == nilai2)
	fmt.Println(nilai2 <= nilai1)
	fmt.Println(nilai1 >= nilai2)
	fmt.Println(nilai2 != nilai1)

}
