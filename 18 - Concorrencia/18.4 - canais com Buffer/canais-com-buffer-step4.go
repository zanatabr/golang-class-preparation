package main

import "fmt"

func main() {

	// [A] Criação do Canal com Buffer
	// (especificando uma capacidade para o canal [2])
	canal := make(chan string, 2)

	canal <- "Olá Mundo"          // [B1] Escrita no canal
	canal <- "Programando em Go!" // [B2] Escrita no canal
	canal <- "Terceira mensagem"  // [B3] Escrita no canal - Deadlock!

	mensagem := <-canal // [C] Espera-se receber algo do Canal

	fmt.Println(mensagem)

	// Aqui ocorre o Deadlock, pois em [B3] estoura a
	// capacidade do buffer, que é de 2 canais.

}
