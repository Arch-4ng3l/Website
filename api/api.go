package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Arch-4ng3l/Website/storage"
)

type APIServer struct {
	listeningAddr string
	storage       storage.Storage
}

func NewAPIServer(listeningAddr string, storage storage.Storage) *APIServer {
	return &APIServer{
		listeningAddr,
		storage,
	}
}

func (s *APIServer) Init() {
	http.HandleFunc("/login", handleLogin)
	http.ListenAndServe(s.listeningAddr, nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		return
	}
	temp, err := ioutil.ReadAll(r.Body)
	body := string(temp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)

}
