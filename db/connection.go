package conection

import (
	"database/sql"
	"fmt"
	"log"
)

type DB struct {
	*sql.DB
}

var DbGlobalUltra2 *DB

func ContectionDB() {
	connStr := "user=postgres password=1000254878 host=localhost port=5432 dbname=co_crud sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	DbGlobalUltra2 = &DB{db}
	fmt.Println("Conection Success")
}
