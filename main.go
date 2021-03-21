package main

import (
	"fmt"

	c "github.com/banco/contas"
)


func main() {

	contaDaSilva := c.ContaCorrente {Titular: "Silvia", Saldo: 300}
	contaDoGustavo := c.ContaCorrente {Titular: "Gustavo", Saldo: 100}

	status := contaDaSilva.Transferir(200, &contaDoGustavo)
	fmt.Println(status)

	fmt.Println(contaDaSilva.Saldo)
	fmt.Println(contaDoGustavo.Saldo)


}