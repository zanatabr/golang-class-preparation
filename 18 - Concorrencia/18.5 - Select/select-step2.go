package main

import (
	"fmt"
	"time"
)

func main() {

	// Criação dos dois canais
	canal1, canal2 := make(chan string), make(chan string)

	// GoRoutine para gerar mensagens destinadas ao canal1
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	// GoRoutine para gerar mensagens destinadas ao canal2
	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	// select: estrutura muito parecida com o "switch", mas é
	// específica para o uso com concorrência
	for {
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}

}
