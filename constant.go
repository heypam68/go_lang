package main

import "fmt"

func main() {
	//single const
	const firstName string = "Muhammad Aulia"
	const lastName = "Rahman"
	const value = 1000

	//fmt.Println(firstName)
	//fmt.Println(lastName)
	//fmt.Println(value)

	//multiple const
	const (
		region    string = "Indonesia"
		age              = 30
		birth            = 12
		namaDepan string = "Muhammad Aulia Rahman"
	)
	fmt.Println(region)
	fmt.Println(age)
	fmt.Println(birth)
	fmt.Println(namaDepan)
}
