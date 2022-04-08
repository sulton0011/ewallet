package api

import (
	"ewallet/api/handlers"
	"ewallet/config"
	"ewallet/storage/repo"

	"github.com/gin-gonic/gin"
)

type Options struct {
	Cfg config.Config
	Repo repo.Repo
}

func New(options Options) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	
	handlers.NewHandler(options.Cfg, options.Repo)

	return router
}