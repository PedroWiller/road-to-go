package main

import (
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
	cachorroJson := `{"nome": "batata", "raca":"vira lata", "idade": 3}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroJson), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorro2Json := `{"nome": "batata", "raca":"vira lata", "idade": 3}`

	c2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorro2Json), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
