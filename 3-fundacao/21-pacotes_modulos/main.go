package main

import (
	"fmt"
	"github.com/valoirmoura/curso-go/3-fundacao/21-pacotes_modulos/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Printf("O Resultado: %d", s)

}
