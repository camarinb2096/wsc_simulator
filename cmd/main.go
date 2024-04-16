package main

import (
	"camarinb2096/wsc_simulator/internal/app/championship"
	"camarinb2096/wsc_simulator/internal/app/matches"
	"camarinb2096/wsc_simulator/internal/app/phases"
	"camarinb2096/wsc_simulator/internal/app/players"
	positions "camarinb2096/wsc_simulator/internal/app/posititons"
	"camarinb2096/wsc_simulator/internal/app/statistics"
	"camarinb2096/wsc_simulator/internal/app/teams"
	"camarinb2096/wsc_simulator/internal/config/db"
	server "camarinb2096/wsc_simulator/internal/config/server"
	logger "camarinb2096/wsc_simulator/pkg"

	"github.com/joho/godotenv"
)

func main() {
	logger := logger.NewLogger()
	logger.Info("Starting WSC-Simulator")

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	mySqlConfig := db.NewDbConfig()
	mySqlDbConn := db.NewDb(mySqlConfig, logger)
	defer db.CloseDb(mySqlDbConn, logger)

	championshipRepo := championship.NewRepository(mySqlDbConn, logger)
	championshipSrv := championship.NewService(championshipRepo, logger)

	matchesRepo := matches.NewRepository(mySqlDbConn, logger)
	matchesSrv := matches.NewService(matchesRepo, logger)

	phasesRepo := phases.NewRepository(mySqlDbConn, logger)
	phasesSrv := phases.NewService(phasesRepo, logger)

	playersRepo := players.NewRepository(mySqlDbConn, logger)
	playersSrv := players.NewService(playersRepo, logger)

	positionsRepo := positions.NewRepository(mySqlDbConn, logger)
	positionsSrv := positions.NewService(positionsRepo, logger)

	statisticsRepo := statistics.NewRepository(mySqlDbConn, logger)
	statisticsSrv := statistics.NewService(statisticsRepo, logger)

	teamsRepo := teams.NewRepository(mySqlDbConn, logger)
	teamsSrv := teams.NewService(teamsRepo, logger)

	server := server.NewServer()
	server.Routes(championshipSrv, matchesSrv, phasesSrv, positionsSrv, statisticsSrv, playersSrv, teamsSrv)
	server.Run(logger)
}
