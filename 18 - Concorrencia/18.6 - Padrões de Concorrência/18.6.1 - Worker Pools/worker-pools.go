package main

import "fmt"

// Worker Pools: tem-se uma fila de tarefas a serem executadas,
//               e vários workers/processos pegando tarefas dessa
//               fila de forma independente para tratá-las.

func main() {

	// canal que terá a sequência de números a serem calculados
	tarefas := make(chan int, 45)

	// canal que armazena os resultados à medida que forem calculados
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	// preenche o canal de tarefas com os valores a serem calculados
	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}

// tarefas (recebe) e resultado (envia)
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
