package server

import (
	"camarinb2096/wsc_simulator/internal/app/championship"
	"camarinb2096/wsc_simulator/internal/app/matches"
	"camarinb2096/wsc_simulator/internal/app/phases"
	"camarinb2096/wsc_simulator/internal/app/players"
	positions "camarinb2096/wsc_simulator/internal/app/posititons"
	"camarinb2096/wsc_simulator/internal/app/statistics"
	"camarinb2096/wsc_simulator/internal/app/teams"
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

func (s *Server) Routes(chmpSrv championship.Services, matchSrv matches.Services, phaseSrv phases.Services, posSrv positions.Services, stcSrv statistics.Services, plaSrv players.Services, teamSrv teams.Services) {
	// chmpEndpoints := championship.NewEndpoints(chmpSrv)
	teamsEndpoints := teams.NewEndpoints(teamSrv)
	playersEndpoints := players.NewEndpoints(plaSrv)
	// matchEndpoints := matches.NewEndpoints(matchSrv)
	// phaseEndpoints := phases.NewEndpoints(phaseSrv)
	// posEndpoints := positions.NewEndpoints(posSrv)
	// stcEndpoints := statistics.NewEndpoints(stcSrv)

	api := s.router.Group("/api/v1")
	{
		api.POST("/teams/upload", gin.HandlerFunc(teamsEndpoints.Upload))
		api.POST("/players/upload", gin.HandlerFunc(playersEndpoints.Upload))
	}

}

func (s *Server) Run(logger *logger.Logger) {
	logger.Info("Starting WSC-Simulator Server")
	s.router.Run()
}
