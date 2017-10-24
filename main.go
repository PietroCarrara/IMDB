package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/apexskier/httpauth"
)

var (
	backend     httpauth.LeveldbAuthBackend
	aaa         httpauth.Authorizer
	port        = 8009
	backendfile = "auth.leveldb"

	db *sql.DB
)

func main() {

	var err error

	db, err = sql.Open("mysql", "IMDB_USER:3T3Dp1uaNXAxbxWv@/IMDB")
	defer db.Close()
	if err != nil {
		log.Print("Erro ao abrir a conexão com o banco em 'main()': " + err.Error())
	}
}
