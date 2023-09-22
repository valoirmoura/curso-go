package main

import "fmt"

func main() {

	var minhaVar interface{} = "Valoir Moura"
	println(minhaVar.(string))

	//nesse caso a variável ok é booleana, ou seja, ela dará true se a corversão for sucesso e false se for o contrário
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)
	//Na conversão acima teremos o valor res=0 e ok = false.... res = 0 que o valor padrão....

}
