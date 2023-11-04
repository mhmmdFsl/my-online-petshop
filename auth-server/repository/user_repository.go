package repository

import (
	"auth-server/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	Save(user *model.User) error
	FindByPrincipal(principal string) (*model.User, error)
}

type userRepositoryImpl struct {
	Db *gorm.DB
}

func (u userRepositoryImpl) FindByPrincipal(principal string) (*model.User, error) {
	var user model.User
	rs := u.Db.Preload(clause.Associations).InnerJoins("PrincipalUser", u.Db.Where(&model.PrincipalUser{PrincipalValue: principal})).First(&user)
	if err := rs.Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (u userRepositoryImpl) Save(user *model.User) error {
	rs := u.Db.Create(user)
	if err := rs.Error; err != nil {
		return err
	}
	rs = u.Db.Save(user)
	if err := rs.Error; err != nil {
		return err
	}
	return nil
}

type UserRepositoryCfg struct {
	Db *gorm.DB
}

func NewUserRepository(c *UserRepositoryCfg) UserRepository {
	return &userRepositoryImpl{
		Db: c.Db,
	}
}
