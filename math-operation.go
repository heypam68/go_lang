package main

import "fmt"

func main() {

	// Penjumlahan langsung tanpa ruang
	var result = 10 * 10

	fmt.Println(result)

	//Penjumlahan dengan ruang
	var a = 20
	var b = 40
	var c = 50
	var d = 100

	var hasil = a + b*c/d

	fmt.Println(hasil)

	//Augmented Operation  1400 -
	var jumlah = 1800

	jumlah -= 400

	fmt.Println(jumlah)

	//Unary Operation
	var i = 200

	i += 100
	i++

	fmt.Println(i)

	// Positive & Negative

	var positive = +100
	var negatif = -100

	fmt.Println(positive)
	fmt.Println(negatif)

}
