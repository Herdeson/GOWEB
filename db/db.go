package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBanco() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=secret@123 host=127.0.0.2 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db

}
