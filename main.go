package main

import (
	_ "blacklist/docs"
	router "blacklist/internal/app/http"
	"blacklist/internal/app/service"
	"blacklist/internal/config"
	"blacklist/internal/repository/postgres"
)

// @title Blacklist
// @version 1.0
// @description Swagger API for Golang Project Blacklist
// @termsOfService http://swagger.io/terms/
// @contact.name Daniil56
// @contact.email daniil13.spiridonov@yandex.ru
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @SecurityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
