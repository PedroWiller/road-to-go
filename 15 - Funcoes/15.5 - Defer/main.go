package main

import "fmt"

func funcao1() {
	fmt.Println("Executando funcao 1")
}

func funcao2() {
	fmt.Println("Executando funcao 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("O resultado da media sera apresentado") //Acredite se quiser.... isso nao executa agora, somente antes de dar o retorno
	fmt.Println("Verificando se aluno aprovado")

	media := (n1 + n2) / 2

	return media > 6
}

func main() {

	defer funcao1()

	//DEFER adia a execucao da funcao

	funcao2()

	fmt.Println(alunoAprovado(7, 8))
}
