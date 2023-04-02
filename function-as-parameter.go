package main

import "fmt"

type seleksiBaju func(string) string

func kumpulanMainan(mainan string, seleksimainan func(string) string) {
	fmt.Println("Ini adalah Mainan", seleksimainan(mainan))
}

func sortirMainan(mainan string) string {
	if mainan == "Lato-lato" {
		return "Yang Sesuai"
	} else {
		return mainan
	}
}

func kataKasar(kata string, seleksiKata func(string) string) {
	fmt.Println("Kata Kata ini Telah Disensor", seleksiKata(kata))
}

func sortirKata(kata string) string {
	if kata == "Anjing" {
		return "..."
	} else {
		return kata
	}
}

func sortirBaju(baju string, seleksiBaju seleksiBaju) {
	fmt.Println("Baju", seleksiBaju(baju))
}

func filterBaju(baju string) string {
	if baju == "Merah" {
		return "Ini Ada Barang"
	} else {
		return baju
	}
}

func main() {
	kumpulanMainan("Remote Control", sortirMainan)
	kumpulanMainan("Lato-lato", sortirMainan)

	kataKasar("Anjing", sortirKata)
	kataKasar("Eko", sortirKata)

	sortirBaju("Merah", filterBaju)
	sortirBaju("-", filterBaju)
}

//Sebenernya ini bisa untuk filtering kata kata dengan function as paramater
