package route

import (
	"auth-server/model"
	"auth-server/model/http_error"
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (r *route) signUp(ctx *fiber.Ctx) error {
	rq := new(model.SignUpRq)
	if err := ctx.BodyParser(&rq); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "FAILED",
			"message": "Failed parse body.",
		})
	}

	if err := r.AuthService.SignUp(rq); err != nil {
		var httpErr http_error.HttpError
		if ok := errors.As(err, &httpErr); ok {
			return ctx.Status(httpErr.StatusCode).JSON(
				fiber.Map{
					"status":  httpErr.Code,
					"message": httpErr.Message,
				})
		}
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "ERROR",
			"message": "Unknown error.",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "SUCCESS",
		"message": "Registration success.",
	})
}
