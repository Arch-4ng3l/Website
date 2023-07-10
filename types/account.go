package types

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Account struct {
	Name     string  `json:"user_name"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
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
	pw := hex.EncodeToString(hash.Sum([]byte(password)))
	fmt.Println("1 " + pw)

	fmt.Println("2 " + a.Password)
	if pw == a.Password {

		return true

	}

	return false
}
