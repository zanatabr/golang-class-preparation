package main

import "fmt"

// Não é muito usado, mas é mais elegante
// retorna uma variável chamada "soma" e outra variável chamada "subtracao"
// os nomes tem efeito somente no escopo da função
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return // Perceber que não é necessário dizer quais são as variáveis de retorno
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}
