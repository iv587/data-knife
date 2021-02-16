package http

import (
	"github.com/gin-gonic/gin"
)

type Route func(engine *gin.Engine)

type Middleware func(engine *gin.Engine)

type Server struct {
	HttpEngine *gin.Engine
	Addr       string
}

func (s *Server) Route(routes Route) {
	routes(s.HttpEngine)
}

func (s *Server) Use(middleware Middleware) {
	middleware(s.HttpEngine)
}

func (s *Server) Start() {
	s.HttpEngine.Run(s.Addr)
}
