package main

import "fmt"

func main() {
	type noKTP string
	type married bool

	var noKtpEko noKTP = "3273151207940002"
	var status married = false

	fmt.Println(noKtpEko)
	fmt.Println(status)

}
