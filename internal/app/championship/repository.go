package championship

import (
	logger "camarinb2096/wsc_simulator/pkg"

	"gorm.io/gorm"
)

type (
	Repository interface {
		GetTeamsId() ([]int, error)
	}

	repo struct {
		db     *gorm.DB
		logger *logger.Logger
	}
)

func NewRepository(db *gorm.DB, logger *logger.Logger) Repository {
	return &repo{
		db:     db,
		logger: logger,
	}
}

func (r *repo) GetTeamsId() ([]int, error) {
	var teamsId []int
	result := r.db.Raw("SELECT id FROM teams").Pluck("id", &teamsId)
	if result.Error != nil {
		r.logger.Error(result.Error.Error())
		return nil, result.Error
	}
	return teamsId, nil
}
