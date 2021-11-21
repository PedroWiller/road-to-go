package servidor

import (
	"crudbasico/banco"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
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

func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	db, erro := banco.Conectar()
	if erro != nil{
		w.Write([]byte("erro ao conectar dados!"))
		return
	}

	defer db.Close()

	query, erro := db.Query("select * from usuarios")
	if erro != nil{
		w.Write([]byte("erro ao criar statement!"))
		return
	}
	defer query.Close()

	var usuarios []usuario
	for query.Next(){
		var usuario usuario
		if erro := query.Scan(&usuario.Id, &usuario.Nome, &usuario.Email);erro !=nil{
			w.Write([]byte("erro ao criar statement!"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if error := json.NewEncoder(w).Encode(usuarios);error != nil{
		w.Write([]byte("erro ao criar statement!"))
		return
	}
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request){
	parametros := mux.Vars(r)
	ID, error := strconv.ParseInt(parametros["id"], 10, 32)
	if error != nil{
		w.Write([]byte("erro ao criar statement!"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil{
		w.Write([]byte("erro ao conectar dados!"))
		return
	}

	defer db.Close()

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil{
		w.Write([]byte("Erro ao buscar usuarios"))
		return
	}

	var usuario usuario
	if linha.Next(){
		if erro:= linha.Scan(&usuario.Id, &usuario.Nome, &usuario.Email);erro!= nil{
			w.Write([]byte("Erro ao buscar usuarios"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario);erro!= nil{
		w.Write([]byte("Erro ao ler usuarios"))
		return
	}
}

func AtualizarUsuarios(w http.ResponseWriter, r *http.Request){
	parametros := mux.Vars(r)
	ID, error := strconv.ParseInt(parametros["id"], 10, 32)
	if error != nil{
		w.Write([]byte("erro ao criar statement!"))
		return
	}

	body, error:= ioutil.ReadAll(r.Body)
	if error !=nil{
		w.Write([]byte("Erro ao ler corpo da request!"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(body, &usuario);erro!=nil{
		w.Write([]byte("Erro ao deserializar corpo"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil{
		w.Write([]byte("erro ao conectar dados!"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro!= nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	if _, erro:= statement.Exec(usuario.Nome, usuario.Email, ID);erro!=nil{
		w.Write([]byte("Erro ao atualizar"))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Delete(w http.ResponseWriter, r *http.Request){
	parametros := mux.Vars(r)
	ID, erro:= strconv.ParseInt(parametros["id"], 10, 32)
	if erro!= nil{
		w.Write([]byte("Erro ao converter parametros"))
		return
	}

	db, erro:= banco.Conectar()
	if erro!= nil{
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	defer db.Close()

	statement, erro:= db.Prepare("Delete from usuarios where id = ?")
	if erro != nil{
		w.Write([]byte("Erro co criar statement"))
		return
	}

	defer statement.Close()

	if _, erro:=statement.Exec(ID);erro!=nil{
		w.Write([]byte("Erro ao deletar usuario"))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sucesso!"))
}
