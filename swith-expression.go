package main

import "fmt"

func main() {
	//Pengecekan kondisi di Switch contoh nama variabel
	name := "Muhammad Aulia Rahman"

	switch name {
	case "Aulia":
		fmt.Println("Nama Benar")
	case "Joko":
		fmt.Println("Nama Benar")
	case "Budi":
		fmt.Println("Nama Benar")
	default:
		fmt.Println("Tidak Ada Kondisi Yang Sesuai dengan Nama")
	}

	//Kondisi Short Statemament di swith
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Terlalu Pendak")
	}

	// Short Statment tanpa kondisi
	length := len(name)
	switch {
	case length < 6:
		fmt.Println("Nama Terlalu Pendek")
	case length < 10:
		fmt.Println("Nama Hampir Sempurna")
	default:
		fmt.Println("Nama Sudah Lengkap")
	}
}
