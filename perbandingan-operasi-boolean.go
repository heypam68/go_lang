package main

import "fmt"

func main() {
	//const lulusGrade bool = "SelamatLulus"
	//Operasi Perbandingan dan &&
	var hargaA = 100000
	var hargaB = 120000
	var hargaC = 170000

	var nilaiBarangA bool = hargaA > 200000
	var nilaiBarangB bool = hargaB < 200000
	var nilaiBarangC bool = hargaC < 200000

	var resultBarang bool = nilaiBarangA && nilaiBarangB && nilaiBarangC

	fmt.Println(resultBarang)

	//Operasi Perbandingan atau ||
	var nilaiAndi = 80
	var nilaiBudi = 90
	var nilaiSiti = 65
	var nilaiAulia = 75

	var lulusKelasAndi bool = nilaiAndi > 85
	var lulusKelasBudi bool = nilaiBudi > 85
	var lulusKelasSiti bool = nilaiSiti > 85
	var lulusKelasAulia bool = nilaiAulia > 85

	var akumulatifRapor bool = lulusKelasAndi || lulusKelasBudi || lulusKelasSiti || lulusKelasAulia

	fmt.Println(akumulatifRapor)

	//Perbandingan Kebalikan

	var nilaiKelulusan = 90
	var nilaiAbsen = 80

	var nilailulus bool = nilaiKelulusan > 65
	var nilaiAbsenLulus bool = nilaiAbsen > 70

	var akumulatifNilaiAbsenLulus = nilailulus && nilaiAbsenLulus

	fmt.Println(akumulatifNilaiAbsenLulus)

}
