package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Arch-4ng3l/Website/types"
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

func (psql *postgresql) FetchUserData(name string) (types.Account, error) {
	query := `SELECT * FROM "users" WHERE "user_name"=$1 VALUES($1)`
	rows, err := psql.DB.Query(query, name)
	if err != nil {
		return types.Account{}, nil
	}
	var id uint
	var user_name string
	var password string
	var balance float64
	rows.Scan(&id, &user_name, &password, &balance)
	acc := types.Account{}

	// Create Account from Fetched Data

	acc.Name = "Name"

	return acc, nil
}
