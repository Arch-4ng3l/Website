package frontend

import (
	"net/http"
)

type Frontend struct {
}

func NewFrontend() *Frontend {
	return &Frontend{}
}

func (f *Frontend) Init() {

	http.HandleFunc("/login", f.handleLogin)
	http.HandleFunc("/signup", f.handleSignUp)
	http.HandleFunc("/", f.handleRequest)

}

func (f *Frontend) handleRequest(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "/home/moritz/Programming/Go/Website/frontend/static/index.html")

}

func (f *Frontend) handleLogin(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "/home/moritz/Programming/Go/Website/frontend/static/login.html")

}

func (f *Frontend) handleSignUp(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "/home/moritz/Programming/Go/Website/frontend/static/signup.html")

}
