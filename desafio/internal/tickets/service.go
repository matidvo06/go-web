package tickets

import (
	"context"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)

type Service struct{
	repo Repository
}

func NewService(repo Repository) *Service{
	return &Service{repo: repo}
}

func (s *Service) GetTotalTickets(ctx context.Context, destination string) (int, error){
	tickets, err := s.repo.GetTicketByDestination(ctx, destination)
	if err != nil{
		return 0, nil
	}
	return len(tickets), nil
}

func (s *Service) AverageDestination(ctx context.Context, destination string) (float64, error){
	tickets, err := s.repo.GetTicketByDestination(ctx, destination)
	if err != nil{
		return 0, err
	}
	totalPeople := 0
	for _, ticket := range tickets{
		totalPeople += ticket.People
	}
	return float64(totalPeople) / float64(len(tickets)), nil
}
  
