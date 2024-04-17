package teams

import (
	"camarinb2096/wsc_simulator/internal/dtos"
	logger "camarinb2096/wsc_simulator/pkg"

	"gorm.io/gorm"
)

type Repository interface {
	Create(team []Team) error
	GetTeamsOrdered() ([]dtos.TeamsOrdered, error)
	GetChampionTeam() (Team, error)
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

func (r *repo) GetTeamsOrdered() ([]dtos.TeamsOrdered, error) {
	var teams []dtos.TeamsOrdered
	result := r.db.Model(&Team{}).Select("id", "name", "points").Order("points desc").Scan(&teams)
	if result.Error != nil {
		r.logger.Error(result.Error.Error())
		return nil, result.Error
	}
	return teams, nil
}

func (r *repo) GetChampionTeam() (Team, error) {
	var team Team
	result := r.db.Model(&Team{}).Select("id", "name", "points").Order("points desc").Limit(1).Scan(&team)
	if result.Error != nil {
		r.logger.Error(result.Error.Error())
		return team, result.Error
	}
	return team, nil
}
