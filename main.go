package main

import (
	"fmt"

	"github.com/Arch-4ng3l/Website/api"
	"github.com/Arch-4ng3l/Website/storage"
)

func main() {

	psql := storage.NewPostgresql("localhost", "moritz", "postgres", "web", 5432)

	server := api.NewAPIServer(":3000", psql)
	fmt.Println("Listening on Port 3000")
	server.Init()
}
