package teams

import (
	logger "camarinb2096/wsc_simulator/pkg"

	"gorm.io/gorm"
)

type Repository interface {
	Create(team []Team) error
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

func (r *repo) Create(team []Team) error {
	err := r.db.Create(&team)
	if err.Error != nil {
		r.logger.Error(err.Error.Error())
		return err.Error
	}
	return nil
}
