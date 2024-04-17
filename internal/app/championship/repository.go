package championship

import (
	logger "camarinb2096/wsc_simulator/pkg"

	"gorm.io/gorm"
)

type (
	Repository interface {
		GetTeamsId() ([]int, error)
		GetQualifiedTeamsByGroup(group []int) ([]int, error)
		RestartPoints() error
		DeleteMatches() error
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

func (r *repo) GetQualifiedTeamsByGroup(group []int) ([]int, error) {
	var teamsId []int
	result := r.db.Raw("SELECT id FROM teams WHERE id IN (?) ORDER BY points DESC LIMIT 2", group).Pluck("id", &teamsId)
	if result.Error != nil {
		r.logger.Error(result.Error.Error())
		return nil, result.Error
	}
	return teamsId, nil
}

func (r *repo) RestartPoints() error {
	result := r.db.Exec("UPDATE teams SET points = 0")
	if result.Error != nil {
		r.logger.Error(result.Error.Error())
		return result.Error
	}
	return nil
}

func (r *repo) DeleteMatches() error {
	result := r.db.Exec("DELETE FROM matches")
	if result.Error != nil {
		r.logger.Error(result.Error.Error())
		return result.Error
	}
	return nil
}
