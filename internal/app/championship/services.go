package championship

import (
	"camarinb2096/wsc_simulator/internal/dtos"
	logger "camarinb2096/wsc_simulator/pkg"
	"fmt"
	"math/rand"
)

type Services interface {
	GroupDraw() dtos.Groups
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

func (s *service) GroupDraw() dtos.Groups {
	teams, _ := s.repo.GetTeamsId()
	fmt.Println(teams)

	if len(teams) < 16 {
		s.logger.Error("There are not enough teams to draw the groups")
		return dtos.Groups{}
	}

	rand.Shuffle(len(teams), func(i, j int) { teams[i], teams[j] = teams[j], teams[i] })

	return dtos.Groups{
		GroupA: []dtos.Group{{FkTeam1: teams[0], FkTeam2: teams[1], FkTeam3: teams[2], FkTeam4: teams[3]}},
		GroupB: []dtos.Group{{FkTeam1: teams[4], FkTeam2: teams[5], FkTeam3: teams[6], FkTeam4: teams[7]}},
		GroupC: []dtos.Group{{FkTeam1: teams[8], FkTeam2: teams[9], FkTeam3: teams[10], FkTeam4: teams[11]}},
		GroupD: []dtos.Group{{FkTeam1: teams[12], FkTeam2: teams[13], FkTeam3: teams[14], FkTeam4: teams[15]}},
	}
}

func (s *service) PlayGroupPhase(dtos.Groups) {

}
