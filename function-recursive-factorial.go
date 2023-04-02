package main

import "fmt"

func factorialLoop(nilai int) int { //Factorial Looping biasa
	result := 1
	for i := nilai; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(nilai int) int { //Factorial Recursive yang ambil nilai nya dari fungsi di atas proses nya
	if nilai == 1 {
		return 1
	} else {
		return nilai * factorialRecursive(nilai-1)
	}
}

func main() {
	loop := factorialLoop(10)                           //Proses Factorial
	fmt.Println(loop)                                   //Outputan Factorial
	fmt.Println(10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1) //Print factorial Biasa

	recursive := factorialRecursive(10)
	fmt.Println(recursive) //Outputan Factorial biasa
}
