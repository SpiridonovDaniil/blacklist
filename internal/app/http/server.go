package http

import (
	"context"

	"blacklist/internal/domain"

	"github.com/gofiber/fiber/v2"
)

//go:generate mockgen -source=server.go -destination=mocks/mock.go

type service interface {
	Create(ctx context.Context, req *domain.Person) error
	Delete(ctx context.Context, req *domain.Id) error
	Get(ctx context.Context, req *domain.Search) ([]*domain.Person, error)
}

func NewServer(service service) *fiber.App {
	f := fiber.New()

	f.Use(HandleErrors)

	f.Post("/", createHandler(service))
	f.Delete("/", deleteHandler(service))
	f.Get("/", getHandler(service))

	return f
}
