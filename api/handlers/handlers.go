package handlers

import (
	"ewallet/api/auth"
	"ewallet/config"
	"ewallet/storage/repo"

	"github.com/go-redis/redis/v8"
)

type Handlers struct {
	cfg   config.Config
	repo  repo.Repo
	auth  auth.Auth
	redis *redis.Client
}

func NewHandler(cfg config.Config, repo repo.Repo, auth auth.Auth, redis *redis.Client) *Handlers {
	return &Handlers{
		cfg:   cfg,
		repo:  repo,
		auth:  auth,
		redis: redis,
	}
}
