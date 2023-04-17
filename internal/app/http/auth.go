package http

import (
	"blacklist/internal/config"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func auth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()
		if err != nil {
			return err
		}
		authHeader := ctx.GetReqHeaders()
		cfg := config.Read()
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
