package main

import (
	"fmt"
	"goPOO/banco/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(1500)
	PagarBoleto(&contaDaLuisa, 1000)

	fmt.Println(contaDaLuisa.ObterSaldo())
}

// func main() {

// 	clienteBruno := clientes.Titular{Nome: "Bruno", CPF: "123.123.123.12", Profissao: "Desenvolvedor"}
// 	contaDoBruno := contas.ContaCorrente{Titular: clienteBruno, NumeroAgencia: 123, NumeroConta: 123456, Saldo: 100}
// 	fmt.Println(contaDoBruno)

// contaExemplo := contas.ContaCorrente{}
// contaExemplo.Depositar(100)

// fmt.Println(contaExemplo)

// contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
//     Nome: "Bruno",
//     CPF: "123.111.123.12",
//     Profissao: "Desenvolvedor"},
//     NumeroAgencia:123, NumeroConta: 123456, Saldo:100}

// fmt.Println(contaDoBruno)

//}

// contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
// contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

// status := contaDoGustavo.Tranferir(0, &contaDaSilvia)

// fmt.Println(status)
// fmt.Println(contaDaSilvia)
// fmt.Println(contaDoGustavo)

// contaDaSilvia := ContaCorrente{}
// contaDaSilvia.titular = "Silvia"
// contaDaSilvia.saldo = 500

// fmt.Println(contaDaSilvia.saldo)

// fmt.Println(contaDaSilvia.Sacar(400)) // tirar duvida com o pit
// fmt.Println(contaDaSilvia.saldo)

// fmt.Println(contaDaSilvia.saldo)
// status, valor := contaDaSilvia.Depositar(2000)
// fmt.Println(status, valor)
