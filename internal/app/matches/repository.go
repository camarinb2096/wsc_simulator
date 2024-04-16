package matches

import (
	logger "camarinb2096/wsc_simulator/pkg"

	"gorm.io/gorm"
)

type Repository interface {
	Create(interface{}) error
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

func (r *repo) Create(interface{}) error {
	return nil
}
