package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Arch-4ng3l/Website/frontend"
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

	http.HandleFunc("/loginAcc", s.handleLogin)

	http.HandleFunc("/signupAcc", s.handleAccountCreate)

	frontend.NewFrontend().Init()

	http.ListenAndServe(s.listeningAddr, nil)

}

func (s *APIServer) handleAccountCreate(w http.ResponseWriter, r *http.Request) {

	var user struct {
		Name string `json:"user_name"`
		Pw   string `json:"password"`
	}

	acc := &types.Account{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	hash := sha256.New()

	acc.Name = user.Name
	acc.Password = hex.EncodeToString(hash.Sum([]byte(user.Pw)))
	acc.Balance = 0

	err = s.storage.CreateNewUser(acc)

	w.WriteHeader(200)
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

	var user struct {
		Name string `json:"user_name"`
		Pw   string `json:"password"`
	}

	json.Unmarshal([]byte(body), &user)

	acc, err := s.storage.FetchUserData(user.Name, user.Pw)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(404)
		return
	}

	if acc == nil {
		w.WriteHeader(400)
		return
	}

	http.Redirect(w, r, "/", 303)

	encode := json.NewEncoder(w)
	encode.Encode(acc)
}
