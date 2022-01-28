package main

import (
	"fmt"
	"time"
)

func main() {

	// aqui na função "main" não foi criada nenhuma Go Routine
	// O Padrão Generator serve para esconder essa complexidade,
	// pois a complexidade de criar a GoRoutine, criação do canal, etc
	//  é encapsulada em uma outra função.

	canal := escrever("Olá Mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// retorna um canal de strings, de uma direção só
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
