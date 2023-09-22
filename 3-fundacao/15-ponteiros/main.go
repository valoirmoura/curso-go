package main

func main() {

	//Memória -> Endereço -> Valor
	//Ponteiro sempre será um endereço de memória

	a := 10
	var ponteiro *int = &a
	*ponteiro = 20

	b := &a
	println(*b)

}
