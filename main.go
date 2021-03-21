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

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) (bool) {
	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo-=valorTransferencia
		status, saldo := contaDestino.Depositar(valorTransferencia)
		fmt.Println(status, saldo)
		return true
	} else {
		return false
	}
}

func main() {

	contaDaSilva := ContaCorrente {titular: "Silvia", saldo: 300}
	contaDoGustavo := ContaCorrente {titular: "Gustavo", saldo: 100}

	status := contaDaSilva.Transferir(200, &contaDoGustavo)
	fmt.Println(status)

	fmt.Println(contaDaSilva.saldo)
	fmt.Println(contaDoGustavo.saldo)


}