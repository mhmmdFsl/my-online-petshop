package service

import (
	"auth-server/model"
	"auth-server/model/http_error"
	"auth-server/repository"
	"auth-server/util"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type AuthService interface {
	SignUp(rq *model.SignUpRq) error
	Login(rq *model.LoginRq) (*model.LoginRs, error)
}

type authServiceImpl struct {
	UserRepo          repository.UserRepository
	PasswordUserRepo  repository.PasswordUserRepository
	PrincipalUserRepo repository.PrincipalUserRepository
	TokenService      TokenService
}

func (a authServiceImpl) Login(rq *model.LoginRq) (*model.LoginRs, error) {
	user, err := a.UserRepo.FindByPrincipal(rq.Principal)
	hashedPassword := []byte(user.PasswordUser.HashPassword)
	password := []byte(rq.Password)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, http_error.NewFailed("User not registered.", http.StatusBadRequest)
		}
		return nil, http_error.NewError(err.Error())
	}
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return nil, http_error.NewFailed("Password doesn't match.", http.StatusBadRequest)
	}
	// generate refresh token
	rt, err := a.TokenService.GenerateRefreshToken(user.Id)
	// generate access token
	token, err := a.TokenService.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}
	// Create the Claims

	return &model.LoginRs{
		User:         user,
		RefreshToken: rt,
		AccessToken:  token,
	}, err
}

func (a authServiceImpl) SignUp(rq *model.SignUpRq) error {

	if ok := util.IsValidEmail(rq.Principal) || util.IsValidPhoneNumber(rq.Principal); !ok {
		return http_error.NewFailed("invalid principal value", http.StatusBadRequest)
	}
	principalType := ""
	if util.IsValidEmail(rq.Principal) {
		principalType = "EMAIL"
	} else {
		principalType = "PHONE"
	}
	exists, err := a.PrincipalUserRepo.ExistsByValue(rq.Principal)
	if err != nil {
		return http_error.NewError(err.Error())
	}

	if exists {
		return http_error.NewFailed("email or phone number already registered", http.StatusBadRequest)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rq.Password), bcrypt.DefaultCost)
	if err != nil {
		return http_error.NewError(err.Error())
	}

	u := model.User{
		Id:        1,
		Name:      rq.Name,
		Status:    "ACTIVE",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		PrincipalUser: model.PrincipalUser{
			PrincipalType:  principalType,
			PrincipalValue: rq.Principal,
			Status:         "ACTIVE",
			CreatedAt:      time.Time{},
			UpdatedAt:      time.Time{},
		},
		PasswordUser: model.PasswordUser{
			HashPassword: string(hashedPassword),
			CreatedAt:    time.Time{},
		},
	}

	err = a.UserRepo.Save(&u)
	if err != nil {
		return http_error.NewError(err.Error())
	}

	return nil
}

type AuthServiceConfig struct {
	UserRepo          repository.UserRepository
	PasswordUserRepo  repository.PasswordUserRepository
	PrincipalUserRepo repository.PrincipalUserRepository
	TokenService      TokenService
}

func NewAuthService(c *AuthServiceConfig) AuthService {
	return &authServiceImpl{
		UserRepo:          c.UserRepo,
		PasswordUserRepo:  c.PasswordUserRepo,
		PrincipalUserRepo: c.PrincipalUserRepo,
		TokenService:      c.TokenService,
	}
}
