package app

import (
	"github.com/gin-gonic/gin"
	"github.io/iv587/goredis-admin/api"
	"github.io/iv587/goredis-admin/cfg"
	"github.io/iv587/goredis-admin/db"
	"github.io/iv587/goredis-admin/http"
)

func Init() {
	err := cfg.InitCfg()
	if err != nil {
		panic(err)
	}
	err = db.Init()
	if err != nil {
		panic(err)
	}
	http.Init()
	gin.SetMode(gin.ReleaseMode)
}

func Start() {
	server := http.Server{
		HttpEngine: gin.Default(),
		Addr:       cfg.GetCfg("server.addr", ":8080"),
	}
	server.Route(api.RouteMap)
	server.Start()
}
