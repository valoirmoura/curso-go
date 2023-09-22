package main

import "fmt"

type Client struct {
	Nome  string
	Idade int
	Ativo bool
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

type Pessoa interface {
	Desativar()
}

func (c Client) Desativar() {
	c.Ativo = false
	fmt.Printf("O Client %s foi desativado\n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	c1 := Client{
		Nome:  "Valoir",
		Idade: 35,
		Ativo: true,
	}

	c1.Ativo = false
	c1.Desativar()

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", c1.Nome, c1.Idade, c1.Ativo)
	Desativacao(c1)

	minhaEmpresa := Empresa{Nome: "Cello Ltda"}
	fmt.Println(minhaEmpresa)
}

//No GO não existe implements por exemplo, Interfaces funcionam assim:
// todos as structs que implementarem o método Desativar automaticamente teram implementado a interface Pessoa
