package main

import "fmt"

func main() {
	//Contoh Break di golang
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}

}
