package main

import "fmt"

// Variática: função que pode receber "n" parâmetros (ou nenhum)
func soma(numeros ...int) int {

	// numeros "é um Slice" de int
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

// É possível ter parâmetros fixos, mas se tiver algum
// parâmetro variável, deve ser declarado na última posição
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 0)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo", 10, 2, 3, 4, 5, 6, 7)
}
