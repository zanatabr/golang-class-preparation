package main

import "fmt"

func main() {

	// OPERADORES ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)
	// FIM DOS OPERADORES ARITMÉTICOS

	// OPERADORES DE ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	// FIM DOS OPERADORES DE ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// FIM DOS OPERADORES RELACIONAIS

	// OPERADORES LÓGICOS
	fmt.Println("--------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	// FIM DOS OPERADORES LÓGICOS

	// OPERADORES UNÁRIOS
	fmt.Println("--------------------------")
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 // numero = numero - 20

	numero *= 3  // numero = numero * 3
	numero /= 10 // numero = numero / 10
	numero %= 3  // numero = numero % 3

	fmt.Println(numero)
	// FIM DOS OPERADORES UNÁRIOS

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)

}
