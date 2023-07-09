package main

import (
	"fmt"

	"github.com/Arch-4ng3l/Website/api"
)

func main() {

	server := api.NewAPIServer(":3000", nil)
	fmt.Println("Listening on Port 3000")
	server.Init()
}
