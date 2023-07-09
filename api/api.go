package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Arch-4ng3l/Website/storage"
	"github.com/Arch-4ng3l/Website/types"
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

	http.HandleFunc("/signup", s.handleAccountCreate)

	http.ListenAndServe(s.listeningAddr, nil)

}

func (s *APIServer) handleAccountCreate(w http.ResponseWriter, r *http.Request) {

	acc := &types.Account{}

	err := json.NewDecoder(r.Body).Decode(acc)

	if err != nil {
		w.WriteHeader(400)
		return
	}
	hash := sha256.New()
	hash.Write([]byte(acc.Password))

	acc.Password = hex.EncodeToString(hash.Sum(nil))

	fmt.Println(acc)
	err = s.storage.CreateNewUser(acc)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(201)
	return
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}
	temp, err := ioutil.ReadAll(r.Body)
	body := string(temp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)

	acc, err := s.storage.FetchUserData(body, "")

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(404)
		return
	}

	encode := json.NewEncoder(w)
	encode.Encode(acc)
}
