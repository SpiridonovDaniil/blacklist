package main

import (
	router "blacklist/internal/app/http"
	"blacklist/internal/app/service"
	"blacklist/internal/config"
	"blacklist/internal/repository/postgres"
)

func main() {
	cfg := config.Read()

	db := postgres.New(cfg.Postgres)
	service := service.New(db)
	r := router.NewServer(service)
	err := r.Listen(":" + cfg.Service.Port)
	if err != nil {
		panic(err)
	}
}
