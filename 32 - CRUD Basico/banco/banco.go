package banco

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Driver conexao
)

func Conectar()(*sql.DB, error) { //entrada - retorno
	stringConexao := "root:golang@/DevBook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil{
		return nil, erro
	}
	if erro = db.Ping(); erro!= nil{
		return nil, erro
	}

	return db, nil
}
