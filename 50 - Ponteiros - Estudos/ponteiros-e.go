package main

import "fmt"

type ponto struct {
	x, y float64
}

func main() {

	// Sintaxe de Ponteiros: &x, px, *px

	// O formato %p mostra o ponteiro em si, não o seu alvo
	// *p : "a coisa" apontada por p (o alvo)

	p := &ponto{2, 3}
	fmt.Printf("p\t%#v\n\n", p)

	var pp *ponto

	pp = new(ponto)
	fmt.Printf("pp\t%#v\n", pp)
	fmt.Printf("\t(%p)\n\n", pp)

	pp = p
	fmt.Printf("pp\t%#v\n", pp)
	fmt.Printf("\t(%p)\n\n", pp)

	fmt.Printf("*pp\t%#v\n\n", *pp)

}
