package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Funcao recuperada com sucesso")
	}
}

func alunoAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()

	media := (n1 + n2) / 2

	if media < 6 {
		return false
	} else if media > 6 {
		return true
	}

	panic("A MEDIA É EXATAMENTE 6")
}

func main() {
	fmt.Println(alunoAprovado(7, 6))
	fmt.Println("Pos Execucao")
}
