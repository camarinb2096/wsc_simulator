package matches

import (
	"camarinb2096/wsc_simulator/internal/dtos"
	logger "camarinb2096/wsc_simulator/pkg"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Create(match Match) error
	GetTotalGoalsByTeam(teamID int) (int, error)
	SumPointToTeam(teamID int, points int) error
	GetMatches() ([]dtos.MatchDetail, error)
	CountMatches() int
	GetMatchStatistics() []dtos.MatchStatistics
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

func (r *repo) GetMatches() ([]dtos.MatchDetail, error) {
	var matches []dtos.MatchDetail
	err := r.db.Raw(`SELECT 
	m.id,
	local.name AS TeamLocalName, 
	visitor.name AS TeamVisitorName,
	m.goals_local,
	m.goals_visitor,
	m.yellow_cards_local,
	m.yellow_cards_visitor,
	m.red_cards_local,
	m.red_cards_visitor,
	winner.name as Winner
  FROM matches m
  INNER JOIN teams local ON m.fk_local_team = local.id
  INNER JOIN teams visitor ON m.fk_visitor_team = visitor.id
  INNER JOIN teams winner ON m.winner = winner.id;
  `).Scan(&matches).Error

	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}
	return matches, nil
}

func (r *repo) CountMatches() int {
	var count int64
	r.db.Model(&Match{}).Count(&count)
	return int(count)
}

func (r *repo) GetMatchStatistics() []dtos.MatchStatistics {
	var statistics []dtos.MatchStatistics
	err := r.db.Raw(`
	SELECT 
    teams.name,
    SUM(goals.goals_scored) AS goals_scored,
    SUM(goals.goals_conceded) AS goals_conceded
FROM (
    SELECT 
        fk_local_team AS team_id, 
        SUM(goals_local) AS goals_scored, 
        SUM(goals_visitor) AS goals_conceded
    FROM matches
    GROUP BY fk_local_team
    UNION ALL
    SELECT 
        fk_visitor_team AS team_id, 
        SUM(goals_visitor) AS goals_scored, 
        SUM(goals_local) AS goals_conceded
    FROM matches
    GROUP BY fk_visitor_team
) AS goals
INNER JOIN teams ON teams.id = goals.team_id
GROUP BY teams.name;
	`).Scan(&statistics).Error

	if err != nil {
		r.logger.Error(err.Error())
		return nil
	}

	return statistics
}
