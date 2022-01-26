package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// waitGroup: serve para sincronizar GoRoutines

	var waitGroup sync.WaitGroup

	waitGroup.Add(4) // Qtde de GoRoutines que serão executadas

	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() // -1 (Tira um cara do waitGroup)
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1 (Tira um cara do waitGroup)
	}()

	go func() {
		escrever("GoRoutine 3!")
		waitGroup.Done() // -1 (Tira um cara do waitGroup)
	}()

	go func() {
		escrever("Goroutine 4!")
		waitGroup.Done() // -1 (Tira um cara do waitGroup)
	}()

	waitGroup.Wait() // 0 (Quando "zerar o group", deixa terminar a execução)
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
