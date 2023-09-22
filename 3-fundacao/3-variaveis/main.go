package main

const a = "Hello, World!"

type ID int //Tipo Customizado, foi Criado um tipo, o tipo ID é um inteiro

var (
	b bool    = true
	c int     = 10
	d string  = "Valoir"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	var nome string = "valoir"
	instr := "Cello"

	println(a, b, c, d, e, nome, instr, f)
}
