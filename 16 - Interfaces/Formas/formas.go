package main

import (
	"fmt"
	"math"
)

// interfaces só contêm assinaturas de métodos
// não possuem campos, como em struct
type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

// Para que retangulo seja considerado uma "forma", deve
// implementar um método chamado "area" que não possua
// parâmetros
// (não faz referência direta a uma interface, diferente do Java)
// A implementação de interfaces é "implícita"
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

// Aqui também: Para que circulo seja considerado uma "forma", deve
// implementar um método chamado "area" que não possua
// parâmetros
// (não faz referência direta a uma interface, diferente do Java)
// A implementação de interfaces é "implícita"
func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {

	// Agora retângulo e circulo satisfazem os
	// requisitos da interface "forma"

	r := retangulo{10, 15}
	escreverArea(r) // A área da forma é 150.00

	c := circulo{10}
	escreverArea(c) // A área da forma é 314.16

}
