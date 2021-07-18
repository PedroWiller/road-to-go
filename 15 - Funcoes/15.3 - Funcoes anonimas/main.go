package main

import "fmt"

func main() {

	retorno := func(texto string) string {
		return fmt.Sprint("Recebido ->", texto)
	}("BATATA")

	fmt.Println(retorno)
}
