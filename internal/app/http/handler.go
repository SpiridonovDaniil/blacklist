package http

import (
	"fmt"
	"net/http"

	_ "blacklist/docs"
	"blacklist/internal/domain"

	"github.com/gofiber/fiber/v2"
)

// Create godoc
// @Summary      create user in blacklist
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        person body domain.Person true "Register person"
// @Success      201
// @Failure      400  {object}  error
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       / [post]
func createHandler(service service) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var req domain.Person
		err := ctx.BodyParser(&req)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return fmt.Errorf("[createHandler] failed to parse request, error: %w", err)
		}
		if checkRequest(req) == true {
			ctx.Status(http.StatusBadRequest)
			return fmt.Errorf("[createHandler] bad request, name, phone number, reason and adding user fields must be filled in")
		}
		err = service.Create(ctx.Context(), &req)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return fmt.Errorf("[createHandler] %w", err)
		}
		ctx.Status(http.StatusCreated)

		return nil
	}
}

func checkRequest(req domain.Person) bool {
	switch "" {
	case req.Name:
		return true
	case req.Phone:
		return true
	case req.Reason:
		return true
	case req.Uploader:
		return true
	}
	return false
}

// Delete godoc
// @Summary      delete user in blacklist
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id body domain.Id true "Delete person"
// @Success      200
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       / [delete]
func deleteHandler(service service) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var req domain.Id
		err := ctx.BodyParser(&req)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return fmt.Errorf("[deleteHandler] failed to parse request, error: %w", err)
		}
		err = service.Delete(ctx.Context(), &req)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return fmt.Errorf("[deleteHandler] %w", err)
		}
		ctx.Status(http.StatusOK)

		return nil
	}
}

// Get godoc
// @Summary      Get user in blacklist
// @Description  get domain.Person by name or phone
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        search body domain.Search true "Show blacklist person"
// @Success      200  {object}  []domain.Person
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /accounts/{id} [get]
func getHandler(service service) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var req domain.Search
		err := ctx.BodyParser(&req)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return fmt.Errorf("[getHandler] failed to parse request, error: %w", err)
		}
		resp, err := service.Get(ctx.Context(), &req)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return fmt.Errorf("[getHandler] %w", err)
		}
		err = ctx.JSON(resp)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return fmt.Errorf("[getHandler] failed to return JSON answer, error: %w", err)
		}

		return nil
	}
}
