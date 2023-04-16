package service

import (
	"context"
	"fmt"
	"time"

	"blacklist/internal/domain"
	"blacklist/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, body *domain.Person) error {
	body.Time = time.Now().Format("02-01-2006")
	err := s.repo.Create(ctx, body)
	if err != nil {
		return fmt.Errorf("[create] error adding a user to the blacklist, error: %w", err)
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, body *domain.Id) error {
	id := body.Id
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("[delete] error deleting a person from the blacklist, error: %w", err)
	}

	return nil
}

func (s *Service) Get(ctx context.Context, body *domain.Search) ([]*domain.Person, error) {
	name := body.Name
	phone := body.Phone
	resp, err := s.repo.Get(ctx, name, phone)
	if err != nil {
		return nil, fmt.Errorf("[get] error searching for a person in the blacklist, error: %w", err)
	}

	return resp, nil
}
