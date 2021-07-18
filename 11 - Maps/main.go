package main

import "fmt"

func main() {
	fmt.Println("Maps")

	//Dentro dos cochetes tipos das chaves e fora tipo dos valores
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobreNome": "Figueredo",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	segundoUsuario := map[string]map[string]string{
		"ususario": {
			"primeiro": "Joao",
			"segundo":  "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}
	fmt.Println(segundoUsuario)

	delete(segundoUsuario, "nome")
	fmt.Println(segundoUsuario)

	segundoUsuario["signo"] = map[string]string{
		"nome": "Escorpião",
	}

	fmt.Println(segundoUsuario)
}
