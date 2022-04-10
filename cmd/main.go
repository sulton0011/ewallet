package main

import (
	"ewallet/api"
	"ewallet/api/auth"
	"ewallet/config"
	"ewallet/storage"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //postgres drivers
)

func main() {
	cfg := config.LoadCfg()
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPass,
		cfg.PostgresDB)

	fmt.Println("port:", cfg.Port)
	connPsql, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		panic(err)
	}

	repo := storage.NewStorage(connPsql).Storage()

	redis := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		DB:       0,
		Password: "",
	})

	api.New(api.Options{
		Cfg:   cfg,
		Repo:  repo,
		Auth:  auth.Auth{Cfg: cfg, Repo: repo},
		Redis: redis,
	})

	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
