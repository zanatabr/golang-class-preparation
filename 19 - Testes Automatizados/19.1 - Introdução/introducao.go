package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {

	// - O teste de uma função nunca deve ficar no mesmo
	//   arquivo da função a ser testada.
	// - O nome do arquivo de teste deve terminar com “_test.go”
	// - Os testes são a única exceção que aceitam dois pacotes
	//   diferentes dentro da mesma pasta

	tipoEndereco := enderecos.TipoDeEndereco("Rodovia dos Imigrantes")
	fmt.Println(tipoEndereco)
}
