package main

import (
	"fmt"
	"time"
)

func main() {

	// Laço que mais se asemelha ao "while" de outras linguagens
	i := 0

	for i < 5 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	fmt.Println("--------------------------")

	// Sintaxe mais parecida com a de outras linguagens
	for j := 0; j < 20; j += 5 {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	fmt.Println("--------------------------")

	// Iterar em algum array/slice/string/map, etc
	nomes := [3]string{"João", "Davi", "Lucas"}

	// O range sempre devolve o índice e o valor
	// O "_" é usado para desprezar o valor
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	fmt.Println("--------------------------")

	// Iterar por uma String
	// - traz o índice e o código da tabela ASCII
	// - Por isso foi usada a função "string()"
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	fmt.Println("--------------------------")

	// Iterar por um Map
	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	fmt.Println("--------------------------")

	// Loop infinito
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}
