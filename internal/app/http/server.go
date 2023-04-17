package http

import (
	"context"

	_ "blacklist/docs"
	"blacklist/internal/config"
	"blacklist/internal/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

//go:generate mockgen -source=server.go -destination=mocks/mock.go

type service interface {
	Create(ctx context.Context, req *domain.Person) error
	Delete(ctx context.Context, req *domain.Id) error
	Get(ctx context.Context, req *domain.Search) ([]*domain.Person, error)
}

func NewServer(service service, cfg *config.Config) *fiber.App {
	f := fiber.New()

	f.Use(HandleErrors)

	f.Get("/swagger/*", swagger.HandlerDefault)

	f.Use(auth(cfg))
	f.Post("/", createHandler(service))
	f.Delete("/", deleteHandler(service))
	f.Get("/", getHandler(service))

	return f
}
