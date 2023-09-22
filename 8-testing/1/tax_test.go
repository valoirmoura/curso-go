package tax

import (
	"testing"
)

// Executando Testes
// Comandos:
// go test,
// go test -coverprofile=coverage.out
// depois de rodar ☝️ podemos "go tool cover -html=coverage.out" para ver onde e qual código precisa ser testado
func TestCalculateTax(t *testing.T) {
	amount := 500
	expected := 5.0

	result := CalculateTax(float64(amount))

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

// Testando mais de uma informação ao mesmo tempo
func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("Expected %f but got %f", item.expected, result)
		}
	}
}

// Benchmark (Comando:
// "go test -bench=." - Esse comando ele roda os testes e também o Benchmark
// "go test -bench=. -run=^#" - Esse comando -run afirma que a gente quer rodar os testes que comecem com ^, como não temos ele não rodara testes... só o benchmark mesmo
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

// Comando para comparar qual algoritimo esta mais performatico...
func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

// Fuzzing
// comando: "go test -fuzz=. -run=^#
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500., 1000.0, 1500.0}

	for _, amout := range seed {
		f.Add(amout)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}
	})
}
