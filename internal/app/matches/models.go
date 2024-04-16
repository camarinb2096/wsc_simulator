package matches

import "gorm.io/gorm"

type Match struct {
	gorm.Model
	FkLocalTeam    uint
	FkVisitorTeam  uint
	FkChampionship uint
	FkPhase        uint
}

func (m *Match) TableName() string {
	return "matches"
}
