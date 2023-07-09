package api

import (
	"encoding/json"
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

	http.HandleFunc("/login", s.handleLogin)
	http.ListenAndServe(s.listeningAddr, nil)
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		return
	}
	temp, err := ioutil.ReadAll(r.Body)
	body := string(temp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)

	// Get DATA
	acc, err := s.storage.FetchUserData(body)

	if err != nil {
		log.Fatal(err)

		// Write An Error to Response
		//w.Write(nil)

		return
	}
	encode := json.NewEncoder(w)
	encode.Encode(acc)

}
