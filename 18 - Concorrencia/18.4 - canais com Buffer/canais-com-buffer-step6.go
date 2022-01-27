package main

import "fmt"

func main() {

	// [A] Criação do Canal com Buffer
	// (especificando uma capacidade para o canal [2])
	canal := make(chan string, 2)

	canal <- "Olá Mundo"          // [B1] Escrita no canal
	canal <- "Programando em Go!" // [B2] Escrita no canal

	mensagem := <-canal  // [C1] Espera-se receber algo do Canal
	mensagem2 := <-canal // [C2] Espera-se receber algo do Canal
	mensagem3 := <-canal // [C3] Espera-se receber algo do Canal - Deadlock

	fmt.Println(mensagem) // Não chegará neste ponto
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)

	// Aqui ocorre o Deadlock em [C3]
}
