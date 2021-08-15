package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Estrutura Select")

	primeiroCanal, segundoCanal := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			primeiroCanal <- "Primeiro Canal"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			segundoCanal <- "Segundo Canal"
		}
	}()

	for {
		select {
		case primeiraMensagem := <-primeiroCanal:
			fmt.Println(primeiraMensagem)
		case segundaMensagem := <-segundoCanal:
			fmt.Println(segundaMensagem)
		}
	}

}
