package product

import (
	"/Users/mdvorsky/otra_carpeta/MODULO4/dia2/ej1t/internal/domain"
	"errors"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	SearchPriceGt(price float64) ([]domain.Product, error)
	Create(p domain.Product) (domain.Product, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]domain.Product, error) {
	l := s.r.GetAll()
	return l, nil
}

func (s *service) GetByID(id int) (domain.Product, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (s *service) SearchPriceGt(price float64) ([]domain.Product, error) {
	l := s.r.SearchPriceGt(price)
	if len(l) == 0 {
		return []domain.Product{}, errors.New("no products found")
	}
	return l, nil
}

func (s *service) Create(p domain.Product) (domain.Product, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}
