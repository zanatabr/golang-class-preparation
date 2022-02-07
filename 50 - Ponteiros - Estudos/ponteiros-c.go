package main

import "fmt"

type ponto struct {
	x int
	y int
}

func main() {

	// Em GO, variáveis são "Caixas"
	// O conteúdo das "Caixas" são independentes

	p := ponto{2, 3}
	p2 := p
	p2.y++

	fmt.Printf("p\t%#v\np2\t%#v\n", p, p2)

}
