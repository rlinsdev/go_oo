package main

import (
	"fmt"

	"github.com/banco/contas"
)

func pagarBoleto (conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)

}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)	
	pagarBoleto(&contaDoDenis, 50)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(50)
	pagarBoleto(&contaDaLuisa, 10)

	fmt.Println(contaDaLuisa.ObterSaldo())
}