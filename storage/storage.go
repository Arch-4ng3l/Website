package storage

import "github.com/Arch-4ng3l/Website/types"

type Storage interface {
	FetchUserData(string, string) (*types.Account, error)
	CreateNewUser(*types.Account) error
}
