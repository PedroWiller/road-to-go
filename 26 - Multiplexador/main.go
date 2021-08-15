package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá mundo"), escrever("Programando em GO"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valores recebido: %s\n", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}

func multiplexar(primeiroCanal, segundoCanal <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-primeiroCanal:
				canalSaida <- mensagem
			case mensagem := <-segundoCanal:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
}
