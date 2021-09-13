package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Load User page!"))
}

func main() {
	//ROTAS
	//URI - Identificador do recurso
	//Metodo - GET, POST, PUT, DELETE
	//URI ==/home
	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
