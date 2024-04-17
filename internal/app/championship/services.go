package championship

import (
	"camarinb2096/wsc_simulator/internal/app/matches"
	"camarinb2096/wsc_simulator/internal/dtos"
	logger "camarinb2096/wsc_simulator/pkg"
	"fmt"
	"math/rand"
)

type Services interface {
	PlayChampionship()
	RestartChampionship() error
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
	qualifiedTeams := PlayGroupPhase(groups, s)
	quarterFinals := PlayQuarterFinals(qualifiedTeams, s)
	semiFinals := PlaySemiFinals(quarterFinals, s)
	winner := PlayFinal(semiFinals, s)

	fmt.Println("Winner: ", winner)

}

func (s *service) RestartChampionship() error {
	err := s.repo.RestartPoints()
	if err != nil {
		s.logger.Error("Error restarting points")
		return err
	}
	err = s.repo.DeleteMatches()
	if err != nil {
		s.logger.Error("Error deleting matches")
		return err
	}
	return nil
}

func GroupDraw(s *service) dtos.Groups {
	teams, _ := s.repo.GetTeamsId()

	if len(teams) < 16 {
		s.logger.Error("There are not enough teams to draw the groups")
		return dtos.Groups{}
	}

	rand.Shuffle(len(teams), func(i, j int) { teams[i], teams[j] = teams[j], teams[i] })

	return dtos.Groups{
		GroupA: teams[0:4],
		GroupB: teams[4:8],
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

func PlayQuarterFinals(qualifieds dtos.Qualifieds, s *service) dtos.Qualifieds {

	quarter1 := s.matchService.PlayMatch(qualifieds.GroupA[0], qualifieds.GroupB[1], 2)
	quarter2 := s.matchService.PlayMatch(qualifieds.GroupA[1], qualifieds.GroupB[0], 2)
	quarter3 := s.matchService.PlayMatch(qualifieds.GroupC[0], qualifieds.GroupD[1], 2)
	quarter4 := s.matchService.PlayMatch(qualifieds.GroupC[1], qualifieds.GroupD[0], 2)

	return dtos.Qualifieds{
		GroupA: []int{quarter1.Winner},
		GroupB: []int{quarter2.Winner},
		GroupC: []int{quarter3.Winner},
		GroupD: []int{quarter4.Winner},
	}
}

func PlaySemiFinals(qualifieds dtos.Qualifieds, s *service) dtos.Qualifieds {
	semi1 := s.matchService.PlayMatch(qualifieds.GroupA[0], qualifieds.GroupB[0], 3)
	semi2 := s.matchService.PlayMatch(qualifieds.GroupC[0], qualifieds.GroupD[0], 3)

	fmt.Println(semi1.Winner, semi2.Winner)

	return dtos.Qualifieds{
		GroupA: []int{semi1.Winner},
		GroupB: []int{semi2.Winner},
	}
}

func PlayFinal(qualifieds dtos.Qualifieds, s *service) int {
	final := s.matchService.PlayMatch(qualifieds.GroupA[0], qualifieds.GroupB[0], 4)

	fmt.Println(final.Winner)

	return final.Winner
}

func playMatchesInGroup(group []int, s *service) {
	s.matchService.PlayMatch(group[0], group[1], 1)
	s.matchService.PlayMatch(group[0], group[2], 1)
	s.matchService.PlayMatch(group[0], group[3], 1)
	s.matchService.PlayMatch(group[1], group[2], 1)
	s.matchService.PlayMatch(group[1], group[3], 1)
	s.matchService.PlayMatch(group[2], group[3], 1)
}
