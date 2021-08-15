package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(4)

	go func() {
		go escrevendo("Salve")
		waitGroup.Done()
	}()

	go func() {
		escrevendo("Programando em GO")
		waitGroup.Done()
	}()

	go func() {
		escrevendo("Go Routine 3")
		waitGroup.Done()
	}()

	go func() {
		escrevendo("Go Routine 4")
		waitGroup.Done()
	}()

	//cada wait group conta -1 quando chegar no wait ele espera o contador chegar ao final da quantidade de go routines
	waitGroup.Wait()
}

func escrevendo(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
