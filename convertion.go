package main

import "fmt"

func main() {
	// Konversi int ke int
	var nilai32 int32 = 3442232
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai8)

	// Konversi tipe data string -> int byte -> string
	var name = "Muhammad"
	var e = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}
