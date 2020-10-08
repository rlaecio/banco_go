package main

import (
	"banco/clientes"
	"fmt"
	"banco/contas"
)


func main() {

	// contaDaCris := contas.ContaCorrente{}
	// contaDaCris.Titular = "Cris"
	// contaDaCris.Saldo = 500
	// contaDaCris.NumeroAgencia = 358
	// contaDaCris.NumeroConta = 123457
	

	// fmt.Println(contaDaCris.Saldo)
	// fmt.Println(contaDaCris.Sacar(300))
	// fmt.Println(contaDaCris.Saldo)
	// fmt.Println(contaDaCris.Depositar(400))

	// // -----------------------
	// contaDaSilvia := contas.ContaCorrente{Titular:"Silvia", Saldo: 300}
    // contaDoGustavo := contas.ContaCorrente{Titular:"Gustavo", Saldo: 100}
    // fmt.Println(contaDaSilvia)
	// fmt.Println(contaDoGustavo)
	
	// status := contaDaSilvia.Transferir(200, &contaDoGustavo)
	// fmt.Println(status)
	// fmt.Println(contaDaSilvia)
	// fmt.Println(contaDoGustavo)
	// contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
	// 	Nome: "Bruno",
	// 	CPF: "123.456.789-10"
	// 	Profissao: "Desenvolvedor"},
	// 	NumeroAgencia: 123, NumeroConta: 123456, Saldo: 100	}
	
	// 	fmt.Println(contaDoBruno)
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo.ObterSaldo())

}
