package main

import (
	"fmt"
)

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string{
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo 
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente)Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito menos que zero", c.saldo
	}
}

func main() {
	contaDaCris := ContaCorrente{}
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500
	contaDaCris.numeroAgencia = 358
	contaDaCris.numeroConta = 123457
	

	fmt.Println(contaDaCris.saldo)
	fmt.Println(contaDaCris.Sacar(300))
	fmt.Println(contaDaCris.saldo)
	fmt.Println(contaDaCris.Depositar(400))

}
