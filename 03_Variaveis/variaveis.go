package main

import "fmt"

func main() {
	var varivel string = "variavel 1"
	variavel2 := "Variavel 2"
	fmt.Println(varivel)
	fmt.Println(variavel2)

	var (
		v1 string = "batata"
		v2 string = "batata tambem somenet que repetido"
	)

	fmt.Println(v1)
	fmt.Println(v2)

	//declaracao inferencia de tipos
	v3, v4 := "v3", "v4"

	fmt.Println(v3)
	fmt.Println(v4)

	const constante string = "GOKU"

	fmt.Println(constante)

	//troca e atribuicao de valores
	v3, v4 = v4, v3

	fmt.Println(v3)
	fmt.Println(v4)
}
