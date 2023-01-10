package http

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Routes func(engine *gin.Engine)

type Middleware func(engine *gin.Engine)

type Server struct {
	HttpEngine *gin.Engine
	Addr       string
}

func (s *Server) Route(routes Routes) {
	routes(s.HttpEngine)
}

func (s *Server) Use(middleware Middleware) {
	middleware(s.HttpEngine)
}

func (s *Server) Start() {
	log.Fatal(s.HttpEngine.Run(s.Addr))
}
