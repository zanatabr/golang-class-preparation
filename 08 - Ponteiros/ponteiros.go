package main

import "fmt"

func main() {

	var variavel1 int = 10
	var variavel2 int = variavel1
	// variavel2 "contém uma cópia" do valor da variavel1
	// ou seja, são distintas

	fmt.Println(variavel1, variavel2)
	// Ex.: 10 10

	variavel2++
	fmt.Println(variavel1, variavel2)
	// perceber que somente o valor da variavel2 foi modificado
	// Ex.: 10 11

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int
	// "ponteiro" guarda o endereço de memória de um valor "int"

	variavel3 = 100
	ponteiro = &variavel3
	// "ponteiro" guarda agora o "endereço de memória" da variavel3

	fmt.Println(variavel3, ponteiro)
	// perceber que foi apresentado o valor contido na variavel3
	// e o endereço de referência armazenado na variável ponteiro
	// Ex.: 100 0xc00001a0e0

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)
	// *ponteiro  ==>  "desreferenciação"
	// perceber que foi apresentado o valor contido na variavel3
	// e o "valor" armazenado no endereço de referência apontado por "ponteiro"
	// Ex.: 100 100

}
