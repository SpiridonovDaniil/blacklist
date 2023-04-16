package repository

import (
	"context"

	"blacklist/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, data *domain.Person) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, name, phone string) ([]*domain.Person, error)
}
