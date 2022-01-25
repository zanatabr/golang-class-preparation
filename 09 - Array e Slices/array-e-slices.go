package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// Array - Declaração explícita
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1) // [Posição 1    ]

	// Array - Declaração por inferência de tipos
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2) // [Posição 1 Posição 2 Posição 3 Posição 4 Posição 5]

	// Array - Declaração: fixa o tamanho do array de acordo com a qtde
	//         de itens passado no momento da declaração
	//         Não torna o Array dinâmico. Não permitiria, por exemplo,
	//         acessar a posição 7 ==>  array3[7] 10  (não permitido)
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3) // [1 2 3 4 5]

	// Slice - Declaração: perceber que não foi definido um tamanho)
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice) // [10 11 12 13 14 15 16 17]

	// Slice (Fatia) não é um Array. Aponta para um Array
	// Para mnostrar que são diferentes:
	fmt.Println(reflect.TypeOf(slice))  // []int
	fmt.Println(reflect.TypeOf(array3)) // [5]int

	// Adiciona item no slice e retorna um slice novo
	slice = append(slice, 18)
	fmt.Println(slice) // [10 11 12 13 14 15 16 17 18]

	// Slice - cria um slice baseado em uma fatia do array2
	slice2 := array2[1:3]
	fmt.Println(slice2) // [Posição 2 Posição 3]

	array2[1] = "Posição Alterada"
	fmt.Println(slice2) // [Posição Alterada Posição 3]
	// Perceber que o slice "apontou" para o conteúdo do array "slice2"

	// ----------------------------------------
	// Arrays Internos
	// ----------------------------------------
	fmt.Println("----------")
	// A função "make" aloca um espaço na memória
	slice3 := make([]float32, 10, 11) // 10 posições, capacidade máxima 11
	fmt.Println(slice3)               // [0 0 0 0 0 0 0 0 0 0]
	// Ou seja, internamente é criado um Array com 11 posições, mas é retornado para
	// o Slice (fatia) apenas as 10 primeiras posições

	slice3 = append(slice3, 5)
	fmt.Println(slice3)      // [0 0 0 0 0 0 0 0 0 0 5]
	fmt.Println(len(slice3)) // length  (foi para 11)
	fmt.Println(cap(slice3)) // capacidade  (continua sendo 11)

	slice3 = append(slice3, 6)
	fmt.Println(slice3)      // [0 0 0 0 0 0 0 0 0 0 5 6]
	fmt.Println(len(slice3)) // length  (foi para 12)
	fmt.Println(cap(slice3)) // capacidade  (AUMENTOU para 24) - Duplicou o tamanho!

	// Perceber que o Slice "nunca" vai estourar o seu tamanho

	// Quando não é informada a capacidade máxima, é assumido que a capacidade = tamanho informado
	slice4 := make([]float32, 5)
	fmt.Println(slice4)      // [0 0 0 0 0]
	fmt.Println(len(slice4)) // 5
	fmt.Println(cap(slice4)) // 5

	slice4 = append(slice4, 10)
	fmt.Println(slice4)      // [0 0 0 0 0 10]
	fmt.Println(len(slice4)) // 6
	fmt.Println(cap(slice4)) // Foi para 12

}
