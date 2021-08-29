package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoEndereco("Rodovia brasil")
	fmt.Println(tipoEndereco)
}
