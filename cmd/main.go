package main

import (
	"ewallet/api"
	"ewallet/config"
	"ewallet/storage"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := config.LoadCfg()
	psqlCred := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPass,
		cfg.PostgresDB,
	)

	connPsql, err := sqlx.Connect("postgres", psqlCred)
	if err != nil {
		panic(err)
	}

	repo := storage.NewStorage(connPsql).Storage()

	server := api.New(api.Options{
		Cfg: cfg,
		Repo: repo,
	})

	if err = server.Run(cfg.Port); err != nil {
		panic(err)
	}
}
