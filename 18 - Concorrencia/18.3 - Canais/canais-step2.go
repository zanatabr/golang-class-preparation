package main

import (
	"fmt"
	"time"
)

func main() {

	canal := make(chan string) // criação do canal. Usa a palavra-chave "chan"

	go escrever("Olá Mundo!", canal) // [A]

	fmt.Println("Depois da função escrever começar a ser executada")

	for {
		mensagem := <-canal   // [B] Aqui o "canal" está aguardando o canal receber um valor
		fmt.Println(mensagem) // [C]
	}

	// DEADLOCK!
	// Aqui foi feito um loop infinito para aguardar todas as possíveis mensagens.
	// Após as 5 iterações, recebemos um:
	//  "falat error: all goroutines are asleep - deadlock!"
	// Informando que não há mais ninguém enviando mensagens para o canal, mas o
	// canal continua aguardando algo a ser recebido, mas isso nunca vai acontecer,
	// o que causa o deadlock.
	// CUIDADO: É um problema que não é percebido em tempo de compilação, apenas em runtime.

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // enviando um valor (texto) para o canal
		time.Sleep(time.Second)
	}
}
