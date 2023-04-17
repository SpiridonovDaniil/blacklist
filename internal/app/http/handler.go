package http

import (
	"fmt"
	"net/http"

	_ "blacklist/docs"
	"blacklist/internal/domain"

	"github.com/gofiber/fiber/v2"
)

// Create godoc
// @Summary      add to blacklist
// @Description  the method adds the user to the blacklist
// @Accept       json
// @Produce      json
// @Param        person body domain.AddPerson true "Name, phone, reason and uploader"
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
// @Summary      remove from blacklist
// @Description  remove a user from the blacklist
// @Accept       json
// @Produce      json
// @Param        id body domain.Id true "User ID"
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
// @Summary      blacklisted search
// @Description  search and get users from the blacklist by phone number or name
// @Accept       json
// @Produce      json
// @Param name query string false "name"
// @Param phone query string false "phone"
// @Success      200  {object}  []domain.Person
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       / [get]
func getHandler(service service) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		name := ctx.Query("name")
		phone := ctx.Query("phone")

		if name == "" && phone == "" {
			ctx.Status(http.StatusBadRequest)
			return fmt.Errorf("[getHandler] search parameters are not specified")
		}
		req := domain.Search{
			Name:  name,
			Phone: phone,
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
