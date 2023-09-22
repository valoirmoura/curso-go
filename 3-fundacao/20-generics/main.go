package main

type Number interface {
	int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	m := map[string]int{"Valoir": 1000, "Edicleia": 2000, "Maria": 3000}
	m2 := map[string]float64{"Valoir": 1000.00, "Edicleia": 2000.40, "Maria": 3000.40}
	println(Soma(m))
	println(Soma(m2))

}
