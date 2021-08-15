package main

import (
	"fmt"
	"linha-comando/app"
	"os"
)

func main() {
	fmt.Println("Started!")

	aplicacao := app.Gerar()
	//os.Args, é OS de entrada do SO
	if erro := aplicacao.Run(os.Args); erro != nil {
		fmt.Println(erro)
	}
}
