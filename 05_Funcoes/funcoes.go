package main

import "fmt"

func Somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := Somar(10, 20)

	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)

		return txt
	}

	resultado := f("Freeza morre")
	fmt.Println(resultado)

	//Funcoes com multiplos retornos
	//detalhe no uso do _, ja usado em outras linguagens, porem para evitar a proxima guerra ele obriga o uso de _ para nao declarar variaveis atoa
	resultadoSoma, _ := calculosMatematicos(2, 15)
	fmt.Println(resultadoSoma)
}
