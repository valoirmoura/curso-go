package main

import "fmt"

type Client struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	c1 := Client{
		Nome:  "Valoir",
		Idade: 35,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", c1.Nome, c1.Idade, c1.Ativo)
	fmt.Println()
	fmt.Println(c1.Nome)

}
