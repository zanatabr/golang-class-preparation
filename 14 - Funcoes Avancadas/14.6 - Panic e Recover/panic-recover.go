package main

import "fmt"

// regra: A média não pode ser 6
// panic: para de executar o que está sendo feito, chama
//   todas as funções que usam "defer". Se não recuperar a
//   função, usando "recover", o programa morre
// Não é um "erro" e não é a melhor alternativa para se trarar erros
func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!")
}

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func main() {
	fmt.Println(alunoEstaAprovado(8, 6))
	fmt.Println("Pós execução!")
}
