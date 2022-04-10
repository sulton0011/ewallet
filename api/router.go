package api

import (
	"ewallet/api/auth"
	"ewallet/api/handlers"
	"ewallet/config"
	"ewallet/storage/repo"
	"io"
	"net/http"

	"github.com/go-redis/redis/v8"
)

type Options struct {
	Cfg   config.Config
	Repo  repo.Repo
	Auth  auth.Auth
	Redis *redis.Client
}

func New(options Options) {
	handler := handlers.NewHandler(options.Cfg, options.Repo, options.Auth, options.Redis)

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			http.HandleFunc("/user/new", handler.AddUser)
			http.HandleFunc("/wallet/new", handler.NewWallet)

			http.HandleFunc("/wallet/check/exist", handler.WalletCheckExists)
			http.HandleFunc("/wallet/check/balance", handler.WalletCheckBalance)
			http.HandleFunc("/wallet/fill", handler.WalletFill)
			http.HandleFunc("/wallet/reduce", handler.WalletReduce)
			http.HandleFunc("/wallet/history", handler.WalletHistory)
		} else {
			io.WriteString(w, "This is a " + r.Method + " request")
		}
	}))
}
