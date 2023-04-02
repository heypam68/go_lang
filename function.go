package main

import "fmt"

func kelengkapanBiodata() {
	biodata := make(map[string]string)
	biodata["Nama"] = "Muhammad Aulia Rahman"
	biodata["Alamat"] = "Jl.Mekar Hurip no 17"
	biodata["Pendidikan Terakhir"] = "D3"

	for key, value := range biodata {
		fmt.Println(key, "=", value)
	}
}

func main() {
	//kelengkapanBiodata()				// Pemanggilan Function
	for i := 0; i < 10; i++ { // Pemanfataan Perulangan yang di panggil functionnya
		kelengkapanBiodata()
	}
}
