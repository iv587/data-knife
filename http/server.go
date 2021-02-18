package http

import (
	"github.com/gin-gonic/gin"
	"github.io/iv587/goredis-admin/cfg"
	"io"
	"log"
	"os"
)

type Route func(engine *gin.Engine)

type Middleware func(engine *gin.Engine)

type Server struct {
	HttpEngine *gin.Engine
	Addr       string
}

func Init() {
	writer := []io.Writer{os.Stdout}
	if cfg.GetBoolCfg("log.enable") {
		f, _ := os.Create(cfg.GetCfg("log.path", "app.log"))
		writer = append(writer, f)
	}
	gin.DefaultWriter = io.MultiWriter(writer...)
}

func (s *Server) Route(routes Route) {
	routes(s.HttpEngine)
}

func (s *Server) Use(middleware Middleware) {
	middleware(s.HttpEngine)
}

func (s *Server) Start() {
	log.Fatal(s.HttpEngine.Run(s.Addr))
}
