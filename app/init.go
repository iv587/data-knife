package app

import (
	"dk/api"
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
		Addr:       ":8080",
	}
	server.Use(api.Cors)
	server.Route(api.Router)
	server.Start()
}

func Start() {
	startDb()
	startHttpServer()
}
