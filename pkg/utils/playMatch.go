package utils

import (
	"camarinb2096/wsc_simulator/internal/dtos"
	"math/rand"
)

func PlayMatch(fkLocalTeam uint, fkVisitorTeam uint, fkPhase uint) dtos.MatchPlayed {
	var matchPlayed dtos.MatchPlayed

	matchPlayed.FkPhase = int(fkPhase)
	matchPlayed.FkLocalTeam = int(fkLocalTeam)
	matchPlayed.FkVisitorTeam = int(fkVisitorTeam)
	matchPlayed.GoalsLocal = 0 + rand.Intn(10-0+1)
	matchPlayed.GoalsVisitor = 0 + rand.Intn(10-0+1)
	matchPlayed.YellowCardsLocal = 0 + rand.Intn(10-0+1)
	matchPlayed.YellowCardsVisitor = 0 + rand.Intn(10-0+1)
	matchPlayed.RedCardsLocal = 0 + rand.Intn(8-0+1)
	matchPlayed.RedCardsVisitor = 0 + rand.Intn(8-0+1)

	return matchPlayed
}
