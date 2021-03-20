package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
	saldo float64
}

func (c * ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 &&  valorDoSaque<= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"		
	}
}

func (c * ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito >0 {
		c.saldo+=valorDoDeposito
		return "Depósito relizado com sucesso", c.saldo
	} else {
		return "Depósito inválido", c.saldo
	}
}

func main() {

	contaDaSilva := ContaCorrente {}
	contaDaSilva.titular = "Silvia"
	contaDaSilva.saldo = 500

	fmt.Println(contaDaSilva.saldo)

	fmt.Println(	contaDaSilva.Depositar(300.64))
}