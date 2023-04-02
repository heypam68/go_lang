package main

import "fmt"

func hello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello" + name
	}
}

func kumpulanBuah(buah string) string {
	if buah == "Mangga" {
		return "Buah Mangga"
	} else {
		return "Buah Kosong" + buah
	}
}

func main() {
	result := hello("Muhammad Aulia Rahman") //Kondisi Salah
	fmt.Println(result)

	fmt.Println(getHello("")) //Kondisi Benar

	resultBuah := kumpulanBuah("Kondisi Salah") //Kondisi Salah
	fmt.Println(resultBuah)

	fmt.Println(kumpulanBuah("Mangga")) //Kondisi Benar
}
