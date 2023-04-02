package main

import "fmt"

func CekNama(nama string) map[string]string {
	if nama == "" {
		return nil
	} else {
		return map[string]string{
			"Ini Adalah": nama,
		}
	}
}

func SortirBarang(barang string) map[string]string {
	if barang == "" {
		return nil
	} else {
		return map[string]string{
			"Barang Ini": barang,
		}
	}
}

func Kelulusan(lulus string) map[string]string {
	if lulus == "" {
		return nil
	} else {
		return map[string]string{
			"Hasil Seleksi dengan Nama": lulus,
		}
	}
}

func Pengecualian(pengeculaian string) map[string]string {
	if pengeculaian == "Tidak lolos Seleksi" {
		return nil
	} else {
		return map[string]string{
			"Hasil Seleksi Anda ": pengeculaian,
		}
	}
}

func main() {
	var nama map[string]string = CekNama("")

	if nama == nil {
		fmt.Println("Nama Tidak ditemukan")
	} else {
		fmt.Println(nama)
	}

	var barang map[string]string = SortirBarang("Toshiba")

	if barang == nil {
		fmt.Println("Kosong")
	} else {
		fmt.Println(barang)
	}

	var lulus map[string]string = Kelulusan("Lulus")

	if lulus == nil {
		fmt.Println("Hasil Anda Belum di Terima")
	} else {
		fmt.Println(lulus)
	}

	var pengecualian map[string]string = Pengecualian("Tidak lolos Seleksi")

	if pengecualian == nil {
		fmt.Println("Hasil Anda Tidak Memenuhi Syarat")
	} else {
		fmt.Println(pengecualian)
	}
}
