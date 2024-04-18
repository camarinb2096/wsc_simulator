package players

import (
	logger "camarinb2096/wsc_simulator/pkg"

	"gorm.io/gorm"
)

type Repository interface {
	Create(players []Player) error
	Get() ([]Player, error)
	GetByTeam(fkTeam int) ([]Player, error)
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

func (r *repo) Get() ([]Player, error) {
	var players []Player
	err := r.db.Find(&players)
	if err.Error != nil {
		r.logger.Error(err.Error.Error())
		return nil, err.Error
	}

	return players, nil
}

func (r *repo) GetByTeam(fkTeam int) ([]Player, error) {
	var players []Player
	err := r.db.Where("fk_team = ?", fkTeam).Find(&players)
	if err.Error != nil {
		r.logger.Error(err.Error.Error())
		return nil, err.Error
	}

	return players, nil
}
