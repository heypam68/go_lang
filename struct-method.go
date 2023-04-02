package main

import "fmt"

type SeleksiCpns struct {
	PendidikanTerakhir string
	IPK                int
	Penempatan         string
	HasilAkhirNilai    int
	HasilAkumulatif    bool
}

type PenerimaanMahasiswa struct {
	NilaiAwal       int
	NilaiAkhir      int
	NilaiAkumulatif int
	HasilAkhir      bool
}

func (penerimaan PenerimaanMahasiswa) penerimaanMahasiswa(HasilAkhir string) {
	fmt.Println("Selamat Lolos Dengan Hasil Akhir", HasilAkhir, "Selamat Anda Lolos", penerimaan.HasilAkhir)
}

func (seleksicpns SeleksiCpns) kumpulanSeleksiCpns(HasilAkumulatif string) {
	fmt.Println("Selamat Hasil Anda", HasilAkumulatif, "Selamat Hasil Anda", seleksicpns.HasilAkumulatif)
}

func (penampungdata SeleksiCpns) ruangData(PendidikanTerakhir string) {
	fmt.Println("Pendidikan Terakhir Anda", PendidikanTerakhir, "Kualifikasi Penidikan Minimal", penampungdata.PendidikanTerakhir)
}

func main() {

	PengumpulanData := SeleksiCpns{
		PendidikanTerakhir: "D3",
		IPK:                3.00,
		Penempatan:         "Cinagara",
		HasilAkhirNilai:    89.00,
		HasilAkumulatif:    true,
	}

	AkumulatifMahasiswa := PenerimaanMahasiswa{
		NilaiAwal:       65,
		NilaiAkhir:      85,
		NilaiAkumulatif: 75,
		HasilAkhir:      true,
	}

	fmt.Println(PengumpulanData)
	fmt.Println(PengumpulanData.PendidikanTerakhir)
	fmt.Println(PengumpulanData.IPK)
	fmt.Println(PengumpulanData.Penempatan)
	fmt.Println(PengumpulanData.HasilAkhirNilai)
	fmt.Println(PengumpulanData.HasilAkumulatif)

	PengumpulanData.kumpulanSeleksiCpns("False")
	PengumpulanData.ruangData("S1")

	fmt.Println(AkumulatifMahasiswa)
	AkumulatifMahasiswa.penerimaanMahasiswa("ini")
}

/**
struct-method adalah memanggil fungsi dari method tersebut
contoh pengumpulan data dan kumpulanseleksicpns
*/
