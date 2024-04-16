package players

import (
	logger "camarinb2096/wsc_simulator/pkg"

	"gorm.io/gorm"
)

type Repository interface {
	Create(players []Player) error
}

type repo struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewRepository(db *gorm.DB, logger *logger.Logger) Repository {
	return &repo{
		db:     db,
		logger: logger,
	}
}

func (r *repo) Create(player []Player) error {
	err := r.db.Create(&player)
	if err.Error != nil {
		r.logger.Error(err.Error.Error())
		return err.Error
	}

	return nil
}
