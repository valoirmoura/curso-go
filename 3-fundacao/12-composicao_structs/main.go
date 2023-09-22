package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {

	c1 := Client{
		Nome:  "Valoir",
		Idade: 35,
		Ativo: true,
	}

	c1.Ativo = false
	c1.Cidade = "Curitiba"
	c1.Numero = 834
	c1.Logradouro = "Rua Rio Parana"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", c1.Nome, c1.Idade, c1.Ativo)
	fmt.Println()
	fmt.Println(c1.Nome)

}
