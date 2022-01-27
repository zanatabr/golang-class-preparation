package main

import "fmt"

func main() {

	// [A] Criação do Canal com Buffer
	// (especificando uma capacidade para o canal [2])
	canal := make(chan string, 2)

	canal <- "Olá Mundo" // [B] Escrita no canal

	mensagem := <-canal // [C] Espera-se receber algo do Canal

	fmt.Println(mensagem)

	// Perceber que agora não ocorreu o Deadlock, pois
	// temos um buffer de 2 canais. Então, só será
	// bloqueado quando atingir a capacidade máxima do
	// buffer.
	//
	// Em [B] enviamos uma mensagem, mas não necessariamente
	// precisamos esperar alguém receber. Isso libera o
	// fluxo do processamento e o programa segue executando.

}
