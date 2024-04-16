package db

import (
	"camarinb2096/wsc_simulator/internal/app/matches"
	"camarinb2096/wsc_simulator/internal/app/phases"
	"camarinb2096/wsc_simulator/internal/app/players"
	positions "camarinb2096/wsc_simulator/internal/app/posititons"
	"camarinb2096/wsc_simulator/internal/app/statistics"
	"camarinb2096/wsc_simulator/internal/app/teams"
	logger "camarinb2096/wsc_simulator/pkg"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewDbConfig() *DbConfig {
	return &DbConfig{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Database: os.Getenv("MYSQL_DATABASE"),
	}
}

func NewDb(cfg *DbConfig, logger *logger.Logger) *gorm.DB {
	logger.Info("Connecting to database")
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})

	if err != nil {
		logger.Fatal("Error connecting to database: %v", err)
	}

	logger.Info("Database connection established")

	return db
}

func CloseDb(db *gorm.DB, logger *logger.Logger) {
	logger.Info("Closing database connection")
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("Error closing database connection: %v", err)
	}

	sqlDB.Close()
}

func Migration(db *gorm.DB, logger *logger.Logger) {
	logger.Info("Running database migrations")
	err := db.AutoMigrate(&teams.Team{}, &players.Player{}, &matches.Match{}, &phases.Phase{}, &positions.Position{}, &statistics.Statistic{})
	if err != nil {
		logger.Fatal("Error running database migrations: %v", err)
	}
	logger.Info("Database migrations completed")
}
