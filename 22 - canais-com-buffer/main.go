package main

import "fmt"

//Para o uso de altecao de valores de canais, é  melhor usar funcoes
//canal com buffer ele nao aguarda alguem enviar em funcao, ele ja da o valor direto
//segundo parametro é a capacidade maxima de uma canal, caso exceder o tamnho do buffer, deadlock!
func main() {
	canal := make(chan string, 2)
	canal <- "Novo teste canal"
	canal <- "Ola mundo"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
