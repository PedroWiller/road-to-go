package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var primeiro int = 10
	var segundo int = primeiro

	fmt.Println(primeiro, segundo)

	primeiro++

	fmt.Println(primeiro, segundo)

	//REFENCIA DE MEMORIAL - ele nao guarda valor, apenas a referencia de memoria de
	var terceiro int = 120
	var ponteiro *int = &terceiro

	fmt.Println(terceiro, *ponteiro)

	terceiro = 100
	ponteiro = &terceiro //& serve para indicar que vai enviar para o endereco

	fmt.Println(terceiro, *ponteiro) //desreferenciacao, remove a referencia, caso contrario aparece o endereco de memoria
}
