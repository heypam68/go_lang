package main

import "fmt"

func main() {
	//Contoh Continue untuk ganjil
	for i := 0; i < 10; i++ {
		if i%2 == 1 { //tergantung kondisi yang di lakukan pada statement misalnya kalo nol biasanya ganjil kalo 1 biasanya genap
			continue
		}
		fmt.Println("Perulangan Ke", i)
	}

	//Contoh Continue untuk Genap
	//for iGenap := 0; iGenap < 10; iGenap++ {
	//	for iGenap%2 == 0 {
	//		continue
	//	}
	//	fmt.Println("Perulangan Ke", iGenap)
	//}

}
