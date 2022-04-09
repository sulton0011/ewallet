package api

import (
	"ewallet/api/handlers"
	"ewallet/config"
	"ewallet/storage/repo"
	"net/http"
)

type Options struct {
	Cfg  config.Config
	Repo repo.Repo
}

func New(options Options) {
	handler := handlers.NewHandler(options.Cfg, options.Repo)

	http.HandleFunc("/", handler.HomePage)
}
