package main

import (
	"fmt"
	"time"
)

func main() {

	canal := make(chan string) // criação do canal. Usa a palavra-chave "chan"

	go escrever("Olá Mundo!", canal) // [A]

	fmt.Println("Depois da função escrever começar a ser executada")

	mensagem := <-canal // [B] Aqui o "canal" está aguardando o canal receber um valor

	fmt.Println(mensagem) // [C]

	// Perceber que aqui foi feita a apresentação de apenas uma mensagem "Olá Mundo!"
	// Mas a func "escrever" possui um laço for com 5 repetições, cadê as outras 4?
	//
	// As operações de "enviar" ou "receber" mensagens para um canal bloqueiam
	// a execução do programa.
	//
	// Neste caso, a execução passa em [A], sem ter a necessidade de aguardar a
	// sua finalização. Quando chega em [B], fica aguardando que algo chegue
	// no "canal" para prosseguir a execução. Quando chegar a primeira mensagem
	// em [B] (não importa se outras mensagens seriam escritas no canal), o
	// fluxo prossegue (desbloqueia) e a execução chega em [C], que apresenta
	// a mensagem na tela. Como é a última instrução a ser executada,
	// finaliza a execução do programa.

}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto // enviando um valor (texto) para o canal
		time.Sleep(time.Second)
	}
}
