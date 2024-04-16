package server

import (
	logger "camarinb2096/wsc_simulator/pkg"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}

func (s *Server) Routes() {
	s.router.GET("/ping")
}

func (s *Server) Run(logger *logger.Logger) {
	logger.Info("Starting WSC-Simulator Server")
	s.router.Run()
}
