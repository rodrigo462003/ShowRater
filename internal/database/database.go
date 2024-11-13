package database

import (
	"database/sql"
	"log"
	"os"

	_"github.com/jackc/pgx/v5"
)


func getConnection() *sql.DB{
	connStr, found := os.LookupEnv("CONN_STR")
	if !found{
		log.Fatalln("Missing environment variable CONN_STR")
	}

	db, err := sql.Open("pgx", connStr)
	if err != nil{
		log.Fatalf("Unable to connect to database because %s", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot ping database because %s", err)
    }
	return db
}
