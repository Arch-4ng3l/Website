package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type postgresql struct {
	DB *sql.DB
}

func NewPostgresql(host, user, password, dbname string, port int) *postgresql {
	psql := postgresql{}
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		host, port, user, password, dbname)
	var err error
	psql.DB, err = sql.Open("postgres", psqlconn)

	if err != nil {
		log.Fatal(err)
	}
	return &psql
}

func (psql *postgresql) fetchUserData(name string) {
	query := `SELECT * FROM "users" WHERE "user_name"=$1 VALUES($1)`
	_, err := psql.DB.Query(query, name)
	if err != nil {
		return
	}
}
