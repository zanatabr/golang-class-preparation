package main

import "fmt"

func main() {

	// [A] Criação do Canal com Buffer
	// (especificando uma capacidade para o canal [2])
	canal := make(chan string, 2)

	canal <- "Olá Mundo"          // [B1] Escrita no canal
	canal <- "Programando em Go!" // [B2] Escrita no canal

	mensagem := <-canal // [C] Espera-se receber algo do Canal

	fmt.Println(mensagem)

	// Aqui também não ocorre o Deadlock, pois
	// temos um buffer de 2 canais.
	//
	// A única falha é que somente uma mensagem [B1] será
	// lida e apresentada, mas segue o fluxo do processamento
	// e o programa segue executando (não trava)

}
