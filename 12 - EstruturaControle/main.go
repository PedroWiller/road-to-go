package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controles")

	numero := -1

	if numero > 15 {
		fmt.Println("Numero maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	//if init
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Segundo numero é maior que zero!")
	} else {
		fmt.Println("Numero é menor que zero!")
	}

}
