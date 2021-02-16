package main

import (
	"github.com/gin-gonic/gin"
	"github.io/iv587/goredis-admin/api"
	"github.io/iv587/goredis-admin/db"
	"github.io/iv587/goredis-admin/http"
	"github.io/iv587/goredis-admin/middleware"
)

func main() {
	db.Init()
	server := http.Server{
		HttpEngine: gin.Default(),
		Addr:       ":8888",
	}
	server.Use(middleware.Middleware)
	server.Route(api.RouteMap)
	server.Start()
}
