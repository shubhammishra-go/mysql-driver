package crud

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Open_DB() (db *sql.DB, err error) {

	db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/")

	log.Println("Database successfully opened!")

	if err != nil {
		panic(err)
	}

	return db, err

}

func Close_DB(db *sql.DB) {

	defer db.Close()

}
