package repository

import (
	"auth-server/model"
	"gorm.io/gorm"
)

type UserRefreshTokenRepository interface {
	Save(userRefreshToken *model.UserRefreshToken) error
	FindByUserId(userId int) (*model.UserRefreshToken, error)
	DeleteByUserId(id int) error
}

type userRefreshTokenRepositoryImpl struct {
	Db *gorm.DB
}

func (u2 userRefreshTokenRepositoryImpl) DeleteByUserId(id int) error {
	var urt model.UserRefreshToken
	err := u2.Db.Where("user_id = ?", id).Delete(&urt).Error
	if err != nil {
		return err
	}

	return nil
}

func (u2 userRefreshTokenRepositoryImpl) Save(u *model.UserRefreshToken) error {
	err := u2.Db.Save(u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u2 userRefreshTokenRepositoryImpl) FindByUserId(userId int) (*model.UserRefreshToken, error) {
	var urt model.UserRefreshToken
	err := u2.Db.Where("user_id", userId).First(&urt).Error
	if err != nil {
		return nil, err
	}
	return &urt, err
}

func NewUserRefreshTokenRepository(db *gorm.DB) UserRefreshTokenRepository {
	return &userRefreshTokenRepositoryImpl{
		Db: db,
	}
}
