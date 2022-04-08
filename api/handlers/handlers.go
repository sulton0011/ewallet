package handlers

import (
	"ewallet/config"
	"ewallet/storage/repo"
)

type handlers struct {
	cfg config.Config
	repo repo.Repo
}

func NewHandler(cfg config.Config, repo repo.Repo) *handlers {
	return &handlers{
		cfg:  cfg,
		repo: repo,
	}
}