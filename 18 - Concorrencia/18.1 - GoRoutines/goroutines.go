package main

import (
	"fmt"
	"time"
)

func main() {

	// Concorrência != Paralelismo
	//
	// Paralelismo:
	// Quando 2 ou mais tarefas executadas ao mesmo tempo.
	// Possível apenas quando se tem mais de um núcleo de CPU

	// Concorrência:
	// As tarefas não estão sendo executadas exatamente ao
	// mesmo tempo. Uma tarefa reveza o tempo de CPU com outra
	// tarefa. Cada uma ocupa um percentual de tempo da CPU

	// ao usar a instrução "go" antes de uma chamada, como no
	// exemplo a seguir, permite que outras chamadas sejam executadas
	// na sequência, sem a necessidade da chamada que foi marcada
	// com "go" terminar. Passa a "concorrer" com as demais chamadas
	go escrever("Olá Mundo!") // goroutine
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
