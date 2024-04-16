package championship

import (
	"camarinb2096/wsc_simulator/internal/app/matches"
	"camarinb2096/wsc_simulator/internal/dtos"
	logger "camarinb2096/wsc_simulator/pkg"
	"math/rand"
)

type Services interface {
	PlayChampionship()
}

type service struct {
	repo         Repository
	logger       *logger.Logger
	matchService matches.Services
}

func NewService(repo Repository, logger *logger.Logger, matchService matches.Services) Services {
	return &service{
		repo:         repo,
		logger:       logger,
		matchService: matchService,
	}
}

func (s *service) PlayChampionship() {
	groups := GroupDraw(s)
	PlayGroupPhase(groups, s)

}

func GroupDraw(s *service) dtos.Groups {
	teams, _ := s.repo.GetTeamsId()

	if len(teams) < 16 {
		s.logger.Error("There are not enough teams to draw the groups")
		return dtos.Groups{}
	}

	rand.Shuffle(len(teams), func(i, j int) { teams[i], teams[j] = teams[j], teams[i] })

	return dtos.Groups{
		GroupA: teams[0:4], // Primeros cuatro equipos para el Grupo A
		GroupB: teams[4:8], // y así sucesivamente...
		GroupC: teams[8:12],
		GroupD: teams[12:16],
	}
}

func PlayGroupPhase(groups dtos.Groups, s *service) dtos.Qualifieds {
	playMatchesInGroup(groups.GroupA, s)
	playMatchesInGroup(groups.GroupB, s)
	playMatchesInGroup(groups.GroupC, s)
	playMatchesInGroup(groups.GroupD, s)

	qualifiedA, _ := s.repo.GetQualifiedTeamsByGroup(groups.GroupA)
	qualifiedB, _ := s.repo.GetQualifiedTeamsByGroup(groups.GroupB)
	qualifiedC, _ := s.repo.GetQualifiedTeamsByGroup(groups.GroupC)
	qualifiedD, _ := s.repo.GetQualifiedTeamsByGroup(groups.GroupD)

	return dtos.Qualifieds{
		GroupA: qualifiedA,
		GroupB: qualifiedB,
		GroupC: qualifiedC,
		GroupD: qualifiedD,
	}
}

func playMatchesInGroup(group []int, s *service) {
	s.matchService.PlayMatch(group[0], group[1], 1)
	s.matchService.PlayMatch(group[0], group[2], 1)
	s.matchService.PlayMatch(group[0], group[3], 1)
	s.matchService.PlayMatch(group[1], group[2], 1)
	s.matchService.PlayMatch(group[1], group[3], 1)
	s.matchService.PlayMatch(group[2], group[3], 1)
}
