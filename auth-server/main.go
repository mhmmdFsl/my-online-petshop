package main

import (
	db2 "auth-server/db"
	"auth-server/repository"
	"auth-server/route"
	"auth-server/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	db, err := db2.NewDb()
	if err != nil {
		log.Fatalf("failed connect to db: %v", err)
		return
	}

	userRepo := repository.NewUserRepository(&repository.UserRepositoryCfg{
		Db: db,
	})

	passwordUserRepo := repository.NewPasswordUserRepository(&repository.PasswordUserRepoCfg{
		Db: db,
	})

	principalUserRepo := repository.NewPrincipalUserRepository(&repository.PrincipalUserRepositoryCfg{
		Db: db,
	})

	userTokenRefreshRepo := repository.NewUserRefreshTokenRepository(db)

	tokenService := service.NewTokenService(&service.TokenServiceCfg{UserRefreshTokenRepo: userTokenRefreshRepo})

	authService := service.NewAuthService(&service.AuthServiceConfig{
		UserRepo:          userRepo,
		PasswordUserRepo:  passwordUserRepo,
		PrincipalUserRepo: principalUserRepo,
		TokenService:      tokenService,
	})

	route.NewRoute(&route.RouterConfig{
		App:         app,
		AuthService: authService,
	})

	err = app.Listen(":3000")
	if err != nil {
		log.Fatal("failed listen to port 3000")
		return
	}
}
