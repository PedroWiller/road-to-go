package main

import "fmt"

func main() {
	//Aritimeticos
	soma := 1 + 2
	subtracao := 22 - 7
	divisao := 21 / 7
	multiplicacao := 23 * 2
	resto := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	//atribuicao
	var v1 string = "BATATA"
	v2 := "BATATA MONSTER"

	fmt.Println(v1, v2)

	//Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//Operadores logicos
	fmt.Println("------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && verdadeiro)
	fmt.Println(falso && verdadeiro)
	fmt.Println(verdadeiro && falso)

	//Operadores Unarios
	numero := 10
	fmt.Println(numero)
	numero++
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero *= 3
	fmt.Println(numero)
	numero %= 4
	fmt.Println(numero)

	//Ternarios
	//texto:= numero > 5 ? "BATATA":"Doido" Non Ecxiste

	var texto string = "BATATA 2"
	if numero > 5 {
		texto = "BATATA"
	}

	fmt.Println(texto)

}
