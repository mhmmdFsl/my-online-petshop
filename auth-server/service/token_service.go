package service

import (
	"auth-server/model"
	"auth-server/model/http_error"
	"auth-server/repository"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenService interface {
	GenerateAccessToken(u *model.User) (string, error)
	GenerateRefreshToken(userId int) (string, error)
}

type tokenServiceImpl struct {
	UserRefreshTokenRepo repository.UserRefreshTokenRepository
}

func (t tokenServiceImpl) GenerateAccessToken(u *model.User) (string, error) {
	claims := jwt.MapClaims{
		"sub": u.Id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	signedString, err := token.SignedString([]byte("S'!L#ld%1;*S:)Mp&<c%1|aZUWBl&GL-"))
	if err != nil {
		return "", http_error.NewError(err.Error())
	}
	return signedString, nil
}

func (t tokenServiceImpl) GenerateRefreshToken(userId int) (string, error) {
	err := t.UserRefreshTokenRepo.DeleteByUserId(userId)
	if err != nil {
		return "", err
	}
	newRefreshToken := utils.UUIDv4()
	urt := &model.UserRefreshToken{
		UserId:       userId,
		RefreshToken: newRefreshToken,
		CreatedAt:    time.Now(),
		ExpiredAt:    time.Now().Add(time.Hour * 24),
	}
	err = t.UserRefreshTokenRepo.Save(urt)
	if err != nil {
		return "", err
	}

	return urt.RefreshToken, nil
}

type TokenServiceCfg struct {
	UserRefreshTokenRepo repository.UserRefreshTokenRepository
}

func NewTokenService(c *TokenServiceCfg) TokenService {
	return &tokenServiceImpl{
		UserRefreshTokenRepo: c.UserRefreshTokenRepo,
	}
}
