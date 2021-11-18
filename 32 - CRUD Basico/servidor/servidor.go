package servidor

import (
	"crudbasico/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	Id    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro!= nil{
		w.Write([]byte("Falha ao ler o corpo da request!"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(requestBody, &usuario); erro!= nil{
		w.Write([]byte("erro ao retornar dados!"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil{
		w.Write([]byte("erro ao conectar dados!"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values(?, ?)")
	if erro != nil{
		w.Write([]byte("erro ao criar statement!"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil{
		w.Write([]byte("erro ao criar statement!"))
		return
	}

	id, erro := insercao.LastInsertId()
	if erro != nil{
		w.Write([]byte("erro ao pegar id inserirdo!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso %d", id)))
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sucesso!"))
}
