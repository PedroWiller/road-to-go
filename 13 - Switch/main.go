package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Nothing"
	}
}

func diaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda"
	case numero == 3:
		return "terça"
	case numero == 4:
		return "Quarta"
	case numero == 5:
		return "Quinta"
	case numero == 6:
		return "Sexta"
	case numero == 7:
		return "Sabado"
	default:
		return "Nothing"
	}
}

func main() {
	fmt.Println("Uso do swithc")

	fmt.Println(diaDaSemana(0))
	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSemana(3))
	fmt.Println(diaDaSemana(2))
	fmt.Println(diaDaSemana(5))

	fmt.Println(diaSemana2(0))
	fmt.Println(diaSemana2(1))
	fmt.Println(diaSemana2(23))
	fmt.Println(diaSemana2(2))
	fmt.Println(diaSemana2(5))
}
