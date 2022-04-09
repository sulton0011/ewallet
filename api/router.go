package api

import (
	"ewallet/api/auth"
	"ewallet/api/handlers"
	"ewallet/config"
	"ewallet/storage/repo"
	"net/http"
)

type Options struct {
	Cfg  config.Config
	Repo repo.Repo
	Auth auth.Auth
}

func New(options Options) {
	handler := handlers.NewHandler(options.Cfg, options.Repo, options.Auth)

	http.HandleFunc("/user", handler.AddUser)
}
