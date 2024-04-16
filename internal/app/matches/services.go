package matches

import (
	logger "camarinb2096/wsc_simulator/pkg"
)

type Services interface {
	CreateMatch() error
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

func (s *service) CreateMatch() error {
	return nil
}
