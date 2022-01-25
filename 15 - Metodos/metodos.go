package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// A função "salvar" é "associada" ao struct "usuário", e torna-se
// um método desse "struct"
// Todos os "usuarios" têm um método "salvar"
func (u usuario) salvar() {
	// Perceber o uso da variável "u"
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// permite modificar um valor de "usuário"
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}
	usuario1.salvar()

	usuario2 := usuario{"Davi", 30}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade) // true

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade) // 31

}
