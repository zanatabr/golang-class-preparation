package main

import (
	"fmt"
	"time"
)

func main() {

	canal := make(chan string) // criação do canal. Usa a palavra-chave "chan"

	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for {
		mensagem, aberto := <-canal // [C] Aqui foi utilizada mais uma var de retorno
		// que identifica se o canal está ou não aberto
		if !aberto { // [D] Se o canal já estiver fechado, interrompe o loop
			break
		}
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!") // [E]

	// [A] Há uma forma mais inteligente de evitar o "deadlock":
	// verificar se o canal está "aberto" ou "fechado", que informa
	// se o canal ainda pode "enviar/receber" dados
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // enviando um valor para o canal
		time.Sleep(time.Second)
	}

	close(canal) // [B] Aqui é feito o fechamento do Canal
}
