package main

import (
	"fmt"

	"github.com/gilsondev/golang-alura-oop/clientes"
	"github.com/gilsondev/golang-alura-oop/contas"
)

func PagarBoleto(conta contas.Conta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {
	clienteSilvia := clientes.Titular{Nome: "Silvia", CPF: "123.123.123-00", Profissao: "Psicologa"}
	clienteGustavo := clientes.Titular{Nome: "Gustavo", CPF: "123.111.123-00", Profissao: "Desenvolvedor"}
	contaDaSilvia := contas.ContaCorrente{Titular: clienteSilvia}
	contaDoGustavo := contas.ContaCorrente{Titular: clienteGustavo}

	contaDaSilvia.Depositar(300)
	contaDoGustavo.Depositar(100)

	status := contaDoGustavo.Transferir(200, &contaDaSilvia)

	PagarBoleto(&contaDaSilvia, 50)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}
