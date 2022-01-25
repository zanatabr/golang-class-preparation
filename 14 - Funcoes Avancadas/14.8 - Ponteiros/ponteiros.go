package main

import "fmt"

// passagem de parâmetro por "valor" (cópia)
func inverterSinal(numero int) int {
	return numero * -1
}

// passagem de parâmetro por "ponteiro" (referência)
func inverterSinalComPonteiro(numero *int) {
	// Não funciona, pois "número" é um endereço de memória
	// numero = numero * -1

	// Aqui é feita a "desreferenciação" de "numero",
	// para manipular o "conteúdo" do endereço passado
	*numero = *numero * -1

	// Não foi necessário usar um "return", pois a alteração está
	// sendo feita diretamente na variável passada
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido) // -20
	fmt.Println(numero)          // 20

	novoNumero := 40
	fmt.Println(novoNumero)               // 40
	inverterSinalComPonteiro(&novoNumero) // Passa o endereço da variável
	fmt.Println(novoNumero)               // -40

}
