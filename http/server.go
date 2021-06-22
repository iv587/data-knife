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
		logPath := cfg.GetCfg("log.path")
		_, err := os.Stat(logPath)
		if err != nil && os.IsNotExist(err) {
			os.MkdirAll(logPath, os.ModePerm)
		}
		f, _ := os.Create(logPath + string(os.PathSeparator) + "http.log")
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
