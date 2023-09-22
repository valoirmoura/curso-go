package main

import "fmt"

func main() {

	total := func() int {
		return sum(1, 3, 4, 5, 6, 7, 8, 9) * 2
	}()

	fmt.Println(total)

}

// esses 3 pontinhos aceita tamanho variatico
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

//Closure é a mesma coisa que função anonima...
