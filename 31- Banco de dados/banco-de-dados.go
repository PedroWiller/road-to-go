package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
//_import implicito
func main() {
	stringConexao := "root:golang@/DevBook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil{
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil{
		log.Fatal(erro)
	}

	fmt.Println("Conexão aberta!")

	linhas, erro := db.Query("SELECT * FROM usuarios")
	if erro!= nil{
		log.Fatal(erro)
	}

	defer linhas.Close()

	fmt.Println(linhas)
}