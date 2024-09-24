package services

import (
	"nugu.dev/rd-vigor/repositories"
)

type PortifolioRepository interface {
	CreatePortifolio(u repositories.User, p repositories.Portifolio) *repositories.RepositoryLayerErr
	GetUserPortifolios(u repositories.User) *repositories.RepositoryLayerErr
}

type PortifolioService struct {
	Repository PortifolioRepository
}

func NewPortifolioService(pr PortifolioRepository) *PortifolioService {
	return &PortifolioService{
		Repository: pr,
	}
}

func (s *PortifolioService) CreatePortifolio(u repositories.User, t string, d string) *ServiceLayerErr {

	if t == "" {
		return &ServiceLayerErr{
			Error:   nil,
			Message: "Error Creating portifolio",
			Code:    400,
		}
	}

	if d == "" {
		return &ServiceLayerErr{
			Error:   nil,
			Message: "Error Creating portifolio",
			Code:    400,
		}
	}

	p := repositories.Portifolio{
		Fk_user_ID:  u.ID,
		Title:       t,
		Description: d,
	}

	if err := s.Repository.CreatePortifolio(u, p); err != nil {
		return &ServiceLayerErr{
			Error:   nil,
			Message: "Error Creating portifolio",
			Code:    500,
		}
	}

	return nil
}

func (s *PortifolioService) GetUserPortifolios(u repositories.User) *ServiceLayerErr {

	if err := s.Repository.GetUserPortifolios(u); err != nil {
		return &ServiceLayerErr{
			Error:   nil,
			Message: "Error Creating portifolio",
			Code:    500,
		}
	}

	return nil
}
