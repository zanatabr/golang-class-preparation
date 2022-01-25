package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := -5

	// if/else simples
	// perceber que não se usa parênteses (apesar de ser permitido)
	// É OBRIGATÓRIO o uso de CHAVES
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// "if init"
	// - inicia variáveis e avalia os valores
	// - as variáveis declaradas e iniciadas aqui não podem ser acessadas
	//   fora do bloco if. Só existem no escopo do bloco if
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}

	// fmt.Println(outroNumero)   // não funcionaria, pois "outroNumero" não existe mais aqui

}
