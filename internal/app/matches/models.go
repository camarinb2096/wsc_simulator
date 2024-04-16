package matches

import "gorm.io/gorm"

type Match struct {
	gorm.Model
	FkLocalTeam        uint
	FkVisitorTeam      uint
	FkChampionship     uint
	FkPhase            uint
	Winner             uint
	GoalsLocal         uint
	GoalsVisitor       uint
	YellowCardsLocal   uint
	YellowCardsVisitor uint
	RedCardsLocal      uint
	RedCardsVisitor    uint
}

func (m *Match) TableName() string {
	return "matches"
}
