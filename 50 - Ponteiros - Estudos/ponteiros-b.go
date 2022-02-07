package main

import "fmt"

func main() {

	// Em GO, variáveis são "Caixas"
	// O conteúdo das "Caixas" são independentes

	a := [...]int{1, 2, 3}
	a2 := a
	a2[0]++

	fmt.Printf("a\t%#v\na2\t%#v\n", a, a2)

}
