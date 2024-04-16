package matches

import (
	logger "camarinb2096/wsc_simulator/pkg"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Create(match Match) error
	GetTotalGoalsByTeam(teamID int) (int, error)
	SumPointToTeam(teamID int, points int) error
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

func (r *repo) Create(match Match) error {
	err := r.db.Model(&Match{}).Create(&match)

	if err.Error != nil {
		r.logger.Error(err.Error.Error())
		return err.Error
	}

	return nil
}

func (r *repo) GetTotalGoalsByTeam(teamID int) (int, error) {
	var totalGoals int
	subQuery := `
		SELECT fk_local_team AS team, IFNULL(goals_local, 0) AS goals
		FROM matches
		WHERE deleted_at IS NULL AND fk_local_team = ?
		UNION ALL
		SELECT fk_visitor_team, IFNULL(goals_visitor, 0)
		FROM matches
		WHERE deleted_at IS NULL AND fk_visitor_team = ?
	`
	err := r.db.Raw(`
		SELECT COALESCE(SUM(goals), 0) AS total_goals
		FROM (`+subQuery+`) AS combined
	`, teamID, teamID).Scan(&totalGoals).Error

	if err != nil {
		return 0, err
	}
	return totalGoals, nil
}

func (r *repo) SumPointToTeam(teamID int, points int) error {
	err := r.db.Exec(`
		UPDATE teams
		SET points = teams.points + ?
		WHERE id = ?
	`, points, teamID).Error

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
