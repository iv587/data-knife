package app

import (
	"dk/api"
	"dk/config"
	"dk/db"
	"dk/http"
	"github.com/gin-gonic/gin"
)

func startDb() {
	db.InitConnection()
}

func startHttpServer() {
	server := http.Server{
		HttpEngine: gin.Default(),
		Addr:       config.GetCfg("server.addr"),
	}
	server.Use(api.Cors)
	server.Route(api.Router)
	server.Start()
}

func Start() {
	config.Init()
	startDb()
	startHttpServer()
}
