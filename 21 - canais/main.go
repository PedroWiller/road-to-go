package main

import (
	"fmt"
	"time"
)

//trem fudido, vc pode usar o paralelismo de processo e setar valores das variaves fora e aguardar a executacao
//apos esse valor dentro da go routines executadas

// o canal tem 2 operacoes. envia e recebe
func main() {
	canal := make(chan string)
	go escrevendo("Salve", canal)

	fmt.Println("Depois da funcao escrever ser executada")

	// for {
	// 	//segunda variavel do canal, pode verificar se ele esta aberto
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}

	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrevendo(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	//ao fim do loop melhor fechar o canal
	close(canal)
}
