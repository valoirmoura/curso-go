package main

import "fmt"

/*
	- Array: possui tamanho fixo, e é possivel iterar sobre ele
*/

func main() {
	var meuArray [3]int //forma de declarar Array, neste caso esse Array possui apenas 3 posições
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(meuArray[len(meuArray)-1])

	for i, v := range meuArray {
		fmt.Printf("O valor do indice é %d é %d\n", i, v)
	}
}

