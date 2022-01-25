package main

import "fmt"

var n int

// Função especial que é executada "antes" da função "main"
// Usada normalmente para fazer um "setup"
// É permitido ter uma função "init" "por arquivo" (não por pacote)
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n) // 10, pois "init" foi executada antes
}
