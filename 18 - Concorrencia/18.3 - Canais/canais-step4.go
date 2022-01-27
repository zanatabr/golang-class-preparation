package main

import (
	"fmt"
	"time"
)

func main() {

	canal := make(chan string) // criação do canal. Usa a palavra-chave "chan"

	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	// Dá pra usar um range para obter cada mensagem recebida pelo
	// canal enquanto ele estiver aberto
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // enviando um valor para o canal
		time.Sleep(time.Second)
	}

	close(canal)
}
