package main

import (
	"crudbasico/servidor"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	_ "github.com/gorilla/mux"
)

func main() {
	//go get github.com/gorilla/mux

	router := mux.NewRouter()

	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/ping", servidor.Ping).Methods(http.MethodGet)

	fmt.Println("Listen port 50002")
	log.Fatal(http.ListenAndServe(":50002", router))
}
