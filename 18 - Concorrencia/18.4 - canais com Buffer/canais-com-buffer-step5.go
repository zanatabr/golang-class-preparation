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

	fmt.Println(mensagem)
	fmt.Println(mensagem2)

	// Aqui não ocorre o Deadlock
}
