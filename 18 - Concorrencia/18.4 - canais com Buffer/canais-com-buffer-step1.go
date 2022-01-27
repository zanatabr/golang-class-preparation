package main

import "fmt"

func main() {

	canal := make(chan string) // [A] Abertura do Canal

	canal <- "Olá Mundo" // [B] Escrita no canal

	mensagem := <-canal // [C] Espera-se receber algo do Canal

	fmt.Println(mensagem)

	// Perceber que ao executar esse programa, recebe-se um:
	// fatal error: all goroutines are asleep - deadlock!
	//
	// Por que?
	// Em [B] temos uma operação bloqueante. Estamos "escrevendo"
	// no canal e esperamos que já tenha alguém pra receber.
	// Quem receberia a mensagem seria o passo [C], mas nunca
	// chegaremos nesse passo. Por isso ocorre esse deadlock
	//
	// É por isso que normalmente os canais são usados em
	// funções separadas, mas há uma alternativa para permanecer
	// na mesma função, que são os canais com buffer

}
