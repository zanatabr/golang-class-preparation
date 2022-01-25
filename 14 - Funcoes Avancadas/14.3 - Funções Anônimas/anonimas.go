package main

import "fmt"

func main() {

	// Declaração e chamada de uma função anônima
	// sem parâmetros e sem retorno
	func() {
		fmt.Println("Teste Funcao Anonima")
	}()

	// Declaração e chamada de uma função anônima
	// com parâmetro e sem retorno
	func(texto string) {
		fmt.Println(texto)
	}("Mais um teste")

	// Declaração e chamada de uma função anônima
	// com parâmetro e cem retorno
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro")

	fmt.Println(retorno)
}
