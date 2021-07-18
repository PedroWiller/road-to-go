package main

import (
	"errors"
	"fmt"
)

func main() {
	//inteiros
	//int8, int16, int32, int64

	//declarar somente int sem tipo de declaracao explicita ele se adequa de acordo com a arquitetura do processador
	var numero int8 = 100
	fmt.Println(numero)

	var numero2 uint32 = 1000 //mesma arquitetura somente aceito positivo
	fmt.Println(numero2)

	//alias

	//rune alias para int32
	var numero3 rune = 123456
	fmt.Println(numero3)

	var numero4 byte = 123 //é adivinha? kkkk
	fmt.Println(numero4)

	//float32 & float64
	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 123.45
	fmt.Println(numeroReal2)

	//valor atribuido por meio da inferencia, o tipo se adapta de acordo com o seu processsador

	//STRING

	var str string = "é mais de 8 mil"
	fmt.Println(str)

	str2 := "Qual o poder de luta de kakaroto?"
	fmt.Println(str2)

	//Declar dessa forma com aspas simples o valor retornado é o valor no tabela ASC
	//Non Existe char no GO, ele somente retorna a Posição
	char := 'B'
	fmt.Println(char)

	//FIM STRING

	//No GO todos os dados tem o valor zero, ou seja, declarou um texto e nao atribui nada ele seta o valor string vazia
	//0 mesmo é valido para inteiros, sempre sera setado o zero
	//Default Values
	var texto string
	fmt.Println(texto)

	//BOOL
	var boolean bool
	fmt.Println(boolean)

	//ERROR
	//errors é o nome do pacote
	//Similar ao notification nosso
	var erro error = errors.New("Erro intero")
	fmt.Println(erro) //valor default nil
}
