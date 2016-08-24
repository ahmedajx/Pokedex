package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func Connect() {
	var err error
	db, err = sql.Open("mysql", "root@/pokedex")
	if err != nil {
		log.Println(err)
	}
}
