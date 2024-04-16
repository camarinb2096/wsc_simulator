package main

import (
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

	server := server.NewServer()
	server.Routes()
	server.Run(logger)
}
