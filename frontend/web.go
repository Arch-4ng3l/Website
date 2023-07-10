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

}

func (f *Frontend) handleLogin(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "/home/moritz/Programming/Go/Website/frontend/static/login.html")

}
