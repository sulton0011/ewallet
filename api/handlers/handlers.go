package handlers

import (
	"ewallet/config"
	"ewallet/storage/repo"
	"fmt"
	"net/http"
)

type Handlers struct {
	cfg  config.Config
	repo repo.Repo
}

func NewHandler(cfg config.Config, repo repo.Repo) *Handlers {
	return &Handlers{
		cfg:  cfg,
		repo: repo,
	}
}

func (h Handlers) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println(r, "Endpoint Hit: homePage")
}
