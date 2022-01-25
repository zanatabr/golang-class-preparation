package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	// 1a forma de criar uma variável Struct
	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	// 2a forma de criar uma variável Struct (por inferência de tipo)
	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	// 2a forma de criar uma variável Struct (por inferência de tipo)
	// Struct composto por outro struct
	usuario2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(usuario2)

	// 3a forma de criar uma variável Struct
	// (por inferência de tipo e definição explícita do campo)
	usuario3 := usuario{idade: 22}
	fmt.Println(usuario3)
}
