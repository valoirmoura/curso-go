package main

import "fmt"

type Client struct {
	nome string
}

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

// Esse é um caso one não utilizar ponteiro de forma alguma
// UseCase: Imagine você fazendo uma simulação de um investimento no seu banco, vc precisa trabalhar com valores simulados,
// ai não utilizamos ponteiro, pois criaremos uma cópia do valor (saldo), no caso saldo atual para trabalhar com as simulações....
func (c Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

// Esse aqui já é uma parada diferente, pois estamos alterando o saldo da conta... aqui precisamos passar um ponteiro, pois o valor será realmente alterado de fato..
func (c *Conta) deposito(valor int) int {
	c.saldo += valor
	return c.saldo
}

func main() {
	conta := Conta{saldo: 100}
	fmt.Println(conta.saldo)
	conta.simular(200)
	println(conta.saldo)

}
