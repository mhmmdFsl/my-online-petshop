package route

import (
	"auth-server/service"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

type route struct {
	App         *fiber.App
	AuthService service.AuthService
}

type RouterConfig struct {
	App         *fiber.App
	AuthService service.AuthService
}

func NewRoute(c *RouterConfig) {

	r := route{
		App:         c.App,
		AuthService: c.AuthService,
	}

	g := r.App.Group("/api/v1/auth")
	g.Post("/sign-up", r.signUp)
	g.Post("/login", r.login)

	r.App.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("wvD5]kc2%Fm}zs~R;mZs6z7d&Ro9Y^cu")},
	}))

	r.App.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hallo")
	})

}
