package main

import (
	"fmt"

	"github.com/bits-and-blooms/bloom/v3"
)

func main() {
	// Estimando parâmetros do Bloom Filter para 1000 elementos com 1% de taxa de falsos positivos
	n := uint(1000)
	fpRate := 0.01
	filter := bloom.NewWithEstimates(n, fpRate)

	// Adicionando elementos
	filter.Add([]byte("apple"))
	filter.Add([]byte("banana"))
	filter.Add([]byte("cherry"))

	// Verificando a presença de elementos
	fmt.Println("Contains 'apple':", filter.Test([]byte("apple")))
	fmt.Println("Contains 'banana':", filter.Test([]byte("banana")))
	fmt.Println("Contains 'cherry':", filter.Test([]byte("cherry")))
	fmt.Println("Contains 'grape':", filter.Test([]byte("grape")))

	// Estimando a taxa de falsos positivos
	m, _ := bloom.EstimateParameters(n, fpRate)
	estimatedFpRate := bloom.EstimateFalsePositiveRate(m, filter.K(), n)
	fmt.Printf("Estimated false positive rate: %f\n", estimatedFpRate)
}
