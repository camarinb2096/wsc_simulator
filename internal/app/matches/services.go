package matches

import (
	"camarinb2096/wsc_simulator/internal/dtos"
	logger "camarinb2096/wsc_simulator/pkg"
	"math/rand"
)

type Services interface {
	PlayMatch(fkLocalTeam int, fkVisitorTeam int, fkPhase int) dtos.Matches
}

type service struct {
	repo   Repository
	logger *logger.Logger
}

func NewService(repo Repository, logger *logger.Logger) Services {
	return &service{
		repo:   repo,
		logger: logger,
	}
}
func (s *service) PlayMatch(fkLocalTeam int, fkVisitorTeam int, fkPhase int) dtos.Matches {
	var matchPlayed dtos.Matches

	matchPlayed.FkPhase = int(fkPhase)
	matchPlayed.FkLocalTeam = int(fkLocalTeam)
	matchPlayed.FkVisitorTeam = int(fkVisitorTeam)
	matchPlayed.GoalsLocal = 0 + rand.Intn(10-0+1)
	matchPlayed.GoalsVisitor = 0 + rand.Intn(10-0+1)
	matchPlayed.YellowCardsLocal = 0 + rand.Intn(10-0+1)
	matchPlayed.YellowCardsVisitor = 0 + rand.Intn(10-0+1)
	matchPlayed.RedCardsLocal = 0 + rand.Intn(8-0+1)
	matchPlayed.RedCardsVisitor = 0 + rand.Intn(8-0+1)

	if matchPlayed.GoalsLocal > matchPlayed.GoalsVisitor {
		matchPlayed.Winner = int(fkLocalTeam)
		s.repo.SumPointToTeam(fkLocalTeam, 3)
	} else if matchPlayed.GoalsLocal < matchPlayed.GoalsVisitor {
		matchPlayed.Winner = int(fkVisitorTeam)
		s.repo.SumPointToTeam(fkVisitorTeam, 3)
	} else {
		totalLocalGoals, _ := s.repo.GetTotalGoalsByTeam(fkLocalTeam)
		totalVisitorGoals, _ := s.repo.GetTotalGoalsByTeam(fkVisitorTeam)
		if totalLocalGoals > totalVisitorGoals {
			matchPlayed.Winner = int(fkLocalTeam)
			s.repo.SumPointToTeam(fkLocalTeam, 3)
		} else {
			matchPlayed.Winner = int(fkVisitorTeam)
			s.repo.SumPointToTeam(fkVisitorTeam, 3)
		}
	}

	s.SavePlayedMatch(matchPlayed)
	return matchPlayed
}

func (s *service) SavePlayedMatch(matchPlayed dtos.Matches) {
	var match Match

	match.FkPhase = uint(matchPlayed.FkPhase)
	match.FkLocalTeam = uint(matchPlayed.FkLocalTeam)
	match.FkVisitorTeam = uint(matchPlayed.FkVisitorTeam)
	match.GoalsLocal = uint(matchPlayed.GoalsLocal)
	match.GoalsVisitor = uint(matchPlayed.GoalsVisitor)
	match.YellowCardsLocal = uint(matchPlayed.YellowCardsLocal)
	match.YellowCardsVisitor = uint(matchPlayed.YellowCardsVisitor)
	match.RedCardsLocal = uint(matchPlayed.RedCardsLocal)
	match.RedCardsVisitor = uint(matchPlayed.RedCardsVisitor)
	match.Winner = uint(matchPlayed.Winner)

	err := s.repo.Create(match)
	if err != nil {
		s.logger.Error(err.Error())
	}
}
