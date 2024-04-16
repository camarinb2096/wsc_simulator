package db

import (
	logger "camarinb2096/wsc_simulator/pkg"
	"fmt"

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
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "password",
		Database: "wsc_simulator",
	}
}

func NewDb(cfg *DbConfig, logger *logger.Logger) *gorm.DB {
	logger.Info("Connecting to database")
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})

	if err != nil {
		logger.Fatal("Error connecting to database: %v", err)
	}

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
