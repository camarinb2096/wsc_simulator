package teams

import (
	logger "camarinb2096/wsc_simulator/pkg"
	"encoding/csv"
	"errors"
	"io"
)

type Services interface {
	Create(file io.Reader) ([]Team, error)
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

func (s *service) Create(file io.Reader) ([]Team, error) {
	s.logger.Info("Reading teams from csv file")

	reader := csv.NewReader(file)
	reader.Comma = ';'
	if _, err := reader.Read(); err != nil {
		return nil, errors.New("error reading header")
	}
	var teams []Team
	count := 0
	for {
		team, err := reader.Read()
		if err == io.EOF {
			break
		}
		if len(team) != 2 {
			s.logger.Error("Error reading teams file")
			return nil, errors.New("error reading teams file")
		}
		if err != nil {
			s.logger.Error("Error reading teams file")
			return nil, errors.New("error reading teams file")
		}
		teams = append(teams, Team{
			Name:      team[0],
			FlagPhoto: team[1],
		})
		count++
	}
	if count < 16 {
		s.logger.Error("insufficient teams in file")
		return nil, errors.New("insufficient teams in file, they must be 16")
	} else if count > 16 {
		s.logger.Error("too many teams in file")
		return nil, errors.New("too many teams in file, they must be 16")
	}

	err := s.repo.Create(teams)
	if err != nil {
		s.logger.Error("Error creating team")
		return nil, errors.New("error creating team")
	}
	return teams, nil
}
