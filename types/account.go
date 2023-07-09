package types

import "crypto/sha256"

type Account struct {
	Name     string `json:"user_name"`
	Password string
	Balance  float64
}

func NewAccount(name, password string, bal float64) *Account {
	return &Account{
		name,
		password,
		bal,
	}
}

func (a *Account) Validate(password string) bool {

	hash := sha256.New()
	hash.Write([]byte(password))

	if string(hash.Sum(nil)) == a.Password {

		return true

	}

	return false
}
