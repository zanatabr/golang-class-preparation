package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// Perceber o mapeamento para JSON nesse strut
type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

// No mapeamento, caso queira "ignorar" algum campo, basta
// mapear com um "hífen". Isso serve para marshal e unmershal
//
// Segue exemplo:
//
// type cachorro struct {
// 	Nome  string `json:"nome"`
// 	Raca  string `json:"-"`
// 	Idade uint   `json:"idade"`
// }
//
// Neste caso, o campo Raca será omitido no JSON

func main() {
	c := cachorro{"Rex", "Dálmata", 3}

	// json.Marshal: Transforma um struct ou map em um JSON
	// usando um struct
	cachorroEmJSON, erro := json.Marshal(c) // Devolve um slice de bytes
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)
	// Mostra: [123 34 110 111 109 101 34 58 34 82 101 120 34 44 34 114
	//          97 99 97 34 58 34 68 195 161 108 109 97 116 97 34 44 34
	//          105 100 97 100 101 34 58 51 125]

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))
	// Mostra: {"nome":"Rex","raca":"Dálmata","idade":3}

	// json.Marshal: Transforma um struct ou map em um JSON
	// Agora usando um map
	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	// json.Marshal: Transforma um struct em um JSON
	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJSON)
	// Mostra: [123 34 110 111 109 101 34 58 34 84 111 98 121 34 44 34
	//          114 97 99 97 34 58 34 80 111 111 100 108 101 34 125]

	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
	// Mostra: {"nome":"Toby","raca":"Poodle"}
}
