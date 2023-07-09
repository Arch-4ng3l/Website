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

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error

	psql.DB, err = sql.Open("postgres", psqlconn)

	if err != nil {
		log.Fatal(err)
	}
	err = psql.DB.Ping()

	if err != nil {
		fmt.Println("DEAD")
		fmt.Println(psql.DB.Ping().Error())
		return nil
	}

	return &psql
}

func (psql *postgresql) CreateNewUser(acc *types.Account) error {

	if psql.DB.Ping() != nil {

		fmt.Println("Connection To DB Died")

	}

	query := `INSERT INTO "users" ("user_name", "password", "account_value")
						VALUES($1, $2, $3);`

	name, password, bal := acc.Name, acc.Password, acc.Balance
	fmt.Println(name + password)
	fmt.Println(bal)
	res, err := psql.DB.Exec(query, name, password, bal)

	if err != nil {

		return err

	}

	n, err := res.RowsAffected()

	if n != 1 || err != nil {
		return fmt.Errorf("Couldnt create User\n")
	}

	return nil

}

func (psql *postgresql) FetchUserData(name, pw string) (*types.Account, error) {

	query := `SELECT * FROM "users" WHERE "user_name"=$1 VALUES($1)`

	rows, err := psql.DB.Query(query, name)

	if err != nil {
		return nil, nil
	}

	var id uint
	var user_name string
	var password string
	var balance float64

	rows.Scan(&id, &user_name, &password, &balance)
	acc := types.NewAccount(name, password, balance)

	if acc.Validate(pw) {
		return acc, nil
	}

	return nil, nil
}
