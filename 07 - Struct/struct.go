package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logadoro string
	numero   int8
}

func main() {
	fmt.Println("Arquivo structs")

	var user usuario

	fmt.Println(user)

	user.idade = 33
	user.nome = "Pedro"

	fmt.Println(user)

	enderecoExemplo := endereco{"ruas dos bobos", 0}
	//Inferencia de tipos
	user2 := usuario{"Pedro", 34, enderecoExemplo}
	fmt.Println(user2)

	user3 := usuario{idade: 35}
	fmt.Println(user3)

}
