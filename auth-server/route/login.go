package route

import (
	"auth-server/model"
	"auth-server/model/http_error"
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (r *route) login(ctx *fiber.Ctx) error {

	var rq = new(model.LoginRq)
	if err := ctx.BodyParser(&rq); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			fiber.Map{
				"status":  "FAILED",
				"message": "Failed parse request body.",
			},
		)
	}

	u, err := r.AuthService.Login(rq)

	if err != nil {
		var httpErr http_error.HttpError
		if ok := errors.As(err, &httpErr); ok {
			return ctx.Status(httpErr.StatusCode).JSON(fiber.Map{
				"status":  httpErr.Code,
				"message": httpErr.Message,
			})
		}
	}

	return ctx.JSON(fiber.Map{
		"status":  "SUCCESS",
		"message": "Login success",
		"data":    u,
	})
}
