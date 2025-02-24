package players

import (
	logger "camarinb2096/wsc_simulator/pkg"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Services interface {
	Create(file io.Reader) error
	Get(fkTeam int) ([]Player, error)
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

func (s *service) Create(file io.Reader) error {
	s.logger.Info("Reading players from csv file")

	reader := csv.NewReader(file)
	reader.Comma = ';'
	if _, err := reader.Read(); err != nil {
		return errors.New("error reading header")
	}
	var players []Player
	count := 0
	for {
		player, err := reader.Read()
		if err == io.EOF {
			break
		}
		if len(player) != 7 {
			fmt.Println("Error reading teams file here!", len(player))
			return errors.New("error reading teams file")
		}
		if err != nil {
			s.logger.Error("Error reading teams file")
			return errors.New("error reading teams file")
		}
		number, err := strconv.Atoi(player[4])
		if err != nil {
			s.logger.Error("Error reading number")
			return errors.New("error reading number")
		}
		fkTeam, err := strconv.Atoi(player[6])
		if err != nil {
			s.logger.Error("Error reading fkTeam")
			return errors.New("error reading fkTeam")
		}
		players = append(players, Player{
			Name:        player[0],
			Nacionality: player[1],
			BirthDate:   player[2],
			Position:    player[3],
			Number:      number,
			Photo:       player[5],
			FkTeam:      fkTeam,
		})
		count++
	}
	err := s.repo.Create(players)
	if err != nil {
		s.logger.Error("Error creating team")
		return errors.New("error creating team")
	}

	return nil
}

func (s *service) Get(fkTeam int) ([]Player, error) {
	var players []Player
	var err error

	if fkTeam == 0 {
		players, err = s.repo.Get()
	} else {
		players, err = s.repo.GetByTeam(fkTeam)
	}
	if err != nil {
		s.logger.Error("Error getting players")
		return nil, errors.New("error getting players")
	}
	return players, nil
}
