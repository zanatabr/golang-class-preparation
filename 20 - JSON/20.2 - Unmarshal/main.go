package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {

	// Unmarshal: transforma um JSON em um struct ou map

	// Conversão para um struct
	// Na conversão é necessário:
	// [A] Criar a variável
	// [B.1] A string que contém o JSON deve ser obrigatoriamente transformada
	//       em um slice de bytes
	// [B.2] É passado o "endereço" (ponteiro) da variável (neste caso c) que
	//       receberá o valor convertido do JSON informado
	cachorroEmJSON := `{"nome":"Rex","raca":"Dálmata","idade":3}`
	var c cachorro                                                       // [A]
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil { // [B]
		log.Fatal(erro)
	}
	fmt.Println(c)
	// Saída: {Rex Dálmata 3}

	// Conversão para um map
	cachorro2EmJSON := `{"nome":"Toby", "raca":"Poodle"}`
	c2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)
	// Saída: map[nome:Toby raca:Poodle]

}
