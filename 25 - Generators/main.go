package main

import (
	"fmt"
	"time"
)

//Padrao generator, acapsula o canal e pode chamar, ele se gernecia
func main() {
	canal := escrever("Olá mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valores recebido: %s\n", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
