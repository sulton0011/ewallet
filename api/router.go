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

	http.HandleFunc("/user/new", handler.AddUser)
	http.HandleFunc("/wallet/new", handler.NewWallet)

	http.HandleFunc("/wallet/check/exist", handler.WalletCheckExists)
	http.HandleFunc("/wallet/check/balance", handler.WalletCheckBalance)
	http.HandleFunc("/wallet/fill", handler.WalletFill)
	http.HandleFunc("/wallet/reduce", handler.WalletReduce)
}
