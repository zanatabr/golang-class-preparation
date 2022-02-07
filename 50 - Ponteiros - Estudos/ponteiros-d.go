package main

import "fmt"

type ponto struct {
	x int
	y int
}

func main() {

	// O operador & (endereço) devolve um ponteiro
	// As caixas "pp" e "pp2" têm ponteiros para a mesma caixa

	pp := &ponto{2, 3}
	pp2 := pp
	pp2.y++

	fmt.Printf("pp\t%#v\npp2\t%#v\n", pp, pp2)

}
