package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int8
	altura    int8
}

//assim se faz heranca, porem se vc tipar como tipo pessoa ela se torna um objeto dentro do outro, sem tipar fica um heranca
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Heranca")

	p1 := pessoa{"joao", "pedro", 28, 100}
	fmt.Println(p1)

	p2 := estudante{p1, "engenharia", "USP"}
	fmt.Println(p2)
}
