package main

import "fmt"

func main() {

	// Toda e qualquer variavel decalarada deve ser usada,
	// se não houver uso para uma variavel declarada, nem mesmo
	// compila/executa o código

	// Declaração explícita (usando var e tipo explícito)
	var variavel1 string = "Variavel 1"

	// Declaração implícita (por inferência) ... perceber o operador ":="
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// Declaração explícita de múltiplas variáveis
	var (
		variavel3 string = "333333"
		variavel4 string = "444444"
	)

	fmt.Println(variavel3, variavel4)

	// Declaração implícita de múltiplas variáveis
	variavel5, variavel6 := "55555555", "6666666"

	fmt.Println(variavel5, variavel6)

	// Constante
	const constante1 string = "Constante 01"
	fmt.Println(constante1)

	const (
		constante2 string = "Constante 02"
		constante3 string = "Constante 03"
	)

	// Troca de valores (não necessita de variável auxiliar)
	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)

}
