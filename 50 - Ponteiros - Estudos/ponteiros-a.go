package main

import "fmt"

func main() {

	// Em GO, variáveis são "Caixas"
	// O conteúdo das "Caixas" são independentes

	i := 3
	i2 := i
	i2++

	fmt.Printf("i\t%#v\ni2\t%#v\n", i, i2)

}
