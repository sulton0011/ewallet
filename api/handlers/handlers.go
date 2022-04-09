package handlers

import (
	"ewallet/api/auth"
	"ewallet/config"
	"ewallet/storage/repo"
)

type Handlers struct {
	cfg  config.Config
	repo repo.Repo
	auth auth.Auth
}

func NewHandler(cfg config.Config, repo repo.Repo, auth auth.Auth) *Handlers {
	return &Handlers{
		cfg:  cfg,
		repo: repo,
		auth: auth,
	}
}
