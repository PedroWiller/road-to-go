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
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/{id}/usuarios", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/ping", servidor.Ping).Methods(http.MethodGet)
	router.HandleFunc("/{id}/usuarios", servidor.AtualizarUsuarios).Methods(http.MethodPut)
	router.HandleFunc("/{id}/usuarios", servidor.Delete).Methods(http.MethodDelete)

	fmt.Println("Listen port 50002")
	log.Fatal(http.ListenAndServe(":50002", router))
}
