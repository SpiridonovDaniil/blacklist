package http

import (
	"fmt"
	"net/http"

	"blacklist/internal/config"

	"github.com/gofiber/fiber/v2"
)

func auth(cfg *config.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()
		if err != nil {
			return err
		}
		authHeader := ctx.GetReqHeaders()
		if len(authHeader["Authorization"]) == 0 {
			ctx.Status(http.StatusUnauthorized)
			return fmt.Errorf("authorization is required Header")
		}
		if authHeader["Authorization"] != cfg.Auth.Auth {
			ctx.Status(http.StatusUnauthorized)
			return fmt.Errorf("this user isn't authorized to this operation: api_key=%s", authHeader["Authorization"])
		}
		return nil
	}
}
