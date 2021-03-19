package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
	saldo float64
}

func main() {

	contaMurldock := ContaCorrente {
		"Murldock",
		589,
		12,
		15.56,
	}

	

	

	fmt.Println(contaMurldock)

}