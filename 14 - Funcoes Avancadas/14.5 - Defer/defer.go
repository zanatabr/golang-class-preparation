package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		// apenas como exemplo: código duplicado,
		// resolvido com o uso do defer
		// fmt.Println("Média calculada. Resultado será retornado!")
		return true
	}
	// fmt.Println("Média calculada. Resultado será retornado!")
	return false
}

func main() {
	// DEFER = Adiar

	// perceber que neste caso, "funcao2" executará "bem" antes da funcao1
	defer funcao1()
	funcao2()
	fmt.Println("-----------------------------")

	fmt.Println(alunoEstaAprovado(7, 8))
	fmt.Println("-----------------------------")
}
