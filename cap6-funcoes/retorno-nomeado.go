package main

import (
	"fmt"
)

func PrecoFinal(precoCusto float64) (precoReal float64, precoDolar float64) {
	fatorLucro := 1.33
	taxaConversao := 2.34

	precoDolar = precoCusto * fatorLucro
	precoReal = precoDolar * taxaConversao

	return precoDolar, precoReal
}

func main() {
	precoDolar, precoReal := PrecoFinal(34.99)

	fmt.Printf("Preço final em dolar: %.2f\n Preço final em real: %.2f\n", precoDolar, precoReal)
}
