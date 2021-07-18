package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println(i)
	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println(j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Joao", "Davi", "Lucas"}

	//se vc tivesse um nome  no index ele ia pedir para usar
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for index, letra := range "BATATA" {
		fmt.Println(index, string(letra))
	}

	usuario := map[string]string{
		"Nome":      "Leonardo",
		"SobreNome": "Batata",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//Não é possivel criar range em struct

	for {
		//isso é um loop infinito
	}
}
