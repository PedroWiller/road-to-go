package main

import (
	"fmt"
	"time"
)

func main() {
	//go routine
	go escrevendo("Salve")
	escrevendo("Programando em GO")
}

func escrevendo(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
