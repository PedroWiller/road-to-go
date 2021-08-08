package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("Teste")
	generica(1)
	generica(true)

	fmt.Println("", 2, true)

	mapa := map[interface{}]interface{}{
		1:            "string",
		float32(100): true,
		"string":     "string",
	}

	fmt.Println(mapa)
}
