package database

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func GetConnectionPool(connectionString string) ( db *sql.DB , err error){
	db, err = sql.Open("pgx", connectionString)
	if err != nil{
		return nil, fmt.Errorf("unable to connect to the database because: %v", err)
	}

	return
}

func UserNameExists(username string, db *sql.DB) (exists bool, err error) {
    username = strings.ToUpper(username)
    err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&exists)
    if err != nil {
        return false, fmt.Errorf("Couldn't query the DB")
    }

    return
}

