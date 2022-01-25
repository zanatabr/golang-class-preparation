package main

import "fmt"

func diaDaSemana(numero int) string {

	// perceber que não se usa "break" como em outras linguagens
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Domingo"
	default:
		return "Número Inválido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	// Outra forma de avaliar os cases
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Domingo"
	default:
		diaDaSemana = "Número Inválido"
	}

	return diaDaSemana
}

func diaDaSemana3(numero int) string {
	var diaDaSemana string

	// Outra forma de avaliar os cases
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough // Não é muito usado. Faz cair no próximo bloco
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Domingo"
	default:
		diaDaSemana = "Número Inválido"
	}

	return diaDaSemana
}
func main() {
	dia := diaDaSemana(200)
	fmt.Println(dia) // Número Inválido

	fmt.Println("-----------")
	dia2 := diaDaSemana2(1)
	fmt.Println(dia2) // "Domingo"

	fmt.Println("-----------")
	dia3 := diaDaSemana3(1)
	fmt.Println(dia3) // "Segunda-Feira"  (percebeu?)
}
