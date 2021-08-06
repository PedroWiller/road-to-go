package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario: %s no banco de dados\n", u.nome)
}

func (u usuario) maiorIdade() bool {
	return u.idade > 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	primeiroUsuario := usuario{"Pedro", 33}
	fmt.Println(primeiroUsuario)
	primeiroUsuario.salvar()

	segundoUsuario := usuario{"Figueredo", 34}
	segundoUsuario.salvar()

	maiorIdade := segundoUsuario.maiorIdade()
	fmt.Println(maiorIdade)

	fmt.Println(segundoUsuario)
	segundoUsuario.fazerAniversario()
	fmt.Println(segundoUsuario)
}
