package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array and Slices")

	var primeiroArray [5]string

	primeiroArray[0] = "BATATA"
	fmt.Println(primeiroArray)

	segundoArray := [5]string{"Vegeta", "Goku", "Bardock"}
	fmt.Println(segundoArray)

	//Isso nao permite ter o tamanho dinamico, ele so permite atribuir o tamanho ja usado
	terceiroArray := [...]int{1, 2, 3, 4, 5}
	fmt.Println(terceiroArray)

	//Slice nao é um array, ele aponta para um array como ponto de referencia
	primeiroSlice := []int{1, 2, 3}
	fmt.Println(primeiroSlice)

	//Cria um novo slice de retorno o append, vc pode adicionar na mesma variavel
	primeiroSlice = append(primeiroSlice, 18)

	fmt.Println(primeiroSlice)

	//O Slice é um pedaco de array, ex:
	segundoSlice := segundoArray[1:3]
	fmt.Println(segundoSlice)

	segundoArray[2] = "KAKAROTO"
	fmt.Println(segundoSlice)

	//Array internos
	//funcao make me lembra o callock
	terceiroSlice := make([]float32, 10, 11)
	fmt.Println(terceiroSlice)
	fmt.Println(len(terceiroSlice))
	fmt.Println(cap(terceiroSlice)) //capacidade

	terceiroSlice = append(terceiroSlice, 3)
	fmt.Println(terceiroSlice)
	fmt.Println(len(terceiroSlice))
	fmt.Println(cap(terceiroSlice)) //capacidade

	terceiroSlice = append(terceiroSlice, 18)
	fmt.Println(terceiroSlice)
	fmt.Println(len(terceiroSlice))
	fmt.Println(cap(terceiroSlice))

	quartoSlice := make([]float32, 5)
	fmt.Println(quartoSlice)

	quartoSlice = append(quartoSlice, 10)

	fmt.Println(len(quartoSlice))
	fmt.Println(cap(quartoSlice))

	//Slice normal javascrit, similar lista .net
}
