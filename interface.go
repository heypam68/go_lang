package main

import "fmt"

type HasName interface { //contract ke getname
	GetName() string
}

type HasAddress interface {
	GetAddress() string
}

type Signature interface {
	Getsign() string
}

type HasAge interface {
	GetAge() int
}

func Umur(umur HasAge) {
	fmt.Println("Umur Anda Saat Ini Adalah", umur.GetAge())
}

func Sign(sign Signature) {
	fmt.Println("Ini Adalah Tanda Tangan", sign.Getsign())
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func SayAnimal(hasAnimal HasName) {
	fmt.Println("Hello", hasAnimal.GetName())
}

func SayAddress(hasaddress HasAddress) {
	fmt.Println("Ini Alamat", hasaddress.GetAddress())
}

type Person struct {
	Name string
}

type Animal struct {
	Animal string
}

type Alamat struct {
	Alamat string
}

type TandaTangan struct {
	TandaTangan string
}

type Kelahiran struct {
	Kelahiran int
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Animal
}

func (alamat Alamat) GetAddress() string {
	return alamat.Alamat
}

func (tandaTangan TandaTangan) Getsign() string {
	return tandaTangan.TandaTangan
}

func (kelahiran Kelahiran) GetAge() int {
	return kelahiran.Kelahiran
}

func main() {
	rahman := Person{
		Name: "Muhammad Aulia Rahman",
	}

	animal := Animal{
		Animal: "Anjing",
	}

	alamat := Alamat{
		Alamat: "Mekar Hurip No 17 kelurahan cijerah kecamatan bandung kulon",
	}

	tandatangan := TandaTangan{
		TandaTangan: "Muhammad Aulia Rahman",
	}

	umur := Kelahiran{
		Kelahiran: 28,
	}

	SayHello(rahman)
	SayAnimal(animal)

	SayAddress(alamat)
	Sign(tandatangan)

	Umur(umur)

}
