package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	primeiroCachorro := cachorro{"Rex", "Dalmata", 12}
	fmt.Println(primeiroCachorro)

	cachorroJSON, erro := json.Marshal(primeiroCachorro)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroJSON)
	fmt.Println(bytes.NewBuffer(cachorroJSON))

	segundoModelo := map[string]string{
		"nome": "toby",
		"raca": "vira lata",
	}

	segundoJSON, erro := json.Marshal(segundoModelo)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(segundoJSON))

}
