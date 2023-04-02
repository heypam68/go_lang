package main

import "fmt"

func perhitunganAngka(angka ...int) int {
	jumlah := 0

	for _, jumlahAngka := range angka {
		jumlah += jumlahAngka
	}
	return jumlah
}

func main() {
	jumlahsemua := perhitunganAngka(10, 20, 40, 50, 60)
	fmt.Println(jumlahsemua)

	slice := []int{20, 30, 40, 50, 100}
	jumlahsemua = perhitunganAngka(slice...)
	fmt.Println(jumlahsemua)
}

//Variadic Function dan slide sama semua bisa mengembalikan data bolak balik
