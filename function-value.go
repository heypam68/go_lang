package main

import "fmt"

func kumpulanNama(nama string) string {
	return "Good Bye" + nama
}

func kumpulanPerhiasan(perhiasan string) string {
	return "Perhiasan adalah" + perhiasan
}

func main() {
	hasilNama := kumpulanNama
	fmt.Println(hasilNama("Muhammad Aulia Rahman"))

	hasilPerhiasan := kumpulanPerhiasan
	result := hasilPerhiasan("Logam")
	fmt.Println(result)

	fmt.Println(kumpulanPerhiasan("Logam Mulia"))
}
