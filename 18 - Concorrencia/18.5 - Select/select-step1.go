package main

import (
	"fmt"
	"time"
)

func main() {

	// Criação dos dois canais
	canal1, canal2 := make(chan string), make(chan string)

	// GoRoutine para gerar mensagens destinadas ao canal1 [A]
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	// GoRoutine para gerar mensagens destinadas ao canal2 [B]
	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	// Aqui temos um problema, pois estamos
	// perdendo tempo e mensagens
	// [A] Gera 1 mensagem a cada 500ms
	// [B] Gera 1 mensagem a cada 2s
	//
	// [C] aguarda 500ms para receber uma mensagem
	// [D] aguarda 2s para receber a mensagem e fica parado
	//     o que vai prejudicar [C] na próxima iteração
	for {
		mensagemCanal1 := <-canal1
		fmt.Println(mensagemCanal1)

		mensagemCanal2 := <-canal2  // [C]
		fmt.Println(mensagemCanal2) // [D]
	}

}
