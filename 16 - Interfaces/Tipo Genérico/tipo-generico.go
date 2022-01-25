package main

import "fmt"

// Como usar uma "interface" para servir como um "tipo genérico"
// Deve ser usado com muito cuidado, para não se tornar uma
// "gambiarra"

// Exemplo cássico: ver implementação da função ftm.Println()

// funcao que recebe uma interface "vazia",
// então serve pra qualquer coisa. Tudo atende (int, boolean, etc)
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String") // String
	generica(1)        // 1
	generica(true)     // true

	// Exemplo cássico: ver implementação da função ftm.Println()
	fmt.Println(1, 2, "string", false, true, float64(12345))

	// Vira uma bagunça, pois a chave pode ser qualquer coisa, e
	// os valores também. Funciona, mas fica uma zona
	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}

	fmt.Println(mapa) // map[true:12 100:true 1:String String:String]
}
