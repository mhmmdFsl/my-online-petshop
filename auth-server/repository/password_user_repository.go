package repository

import (
	"auth-server/model"
	"gorm.io/gorm"
)

type PasswordUserRepository interface {
	Save(pu *model.PasswordUser) error
}

type passwordUserRepoImpl struct {
	Db *gorm.DB
}

func (p passwordUserRepoImpl) Save(pu *model.PasswordUser) error {
	rs := p.Db.Save(pu)
	if err := rs.Error; err != nil {
		return err
	}

	return nil
}

type PasswordUserRepoCfg struct {
	Db *gorm.DB
}

func NewPasswordUserRepository(c *PasswordUserRepoCfg) PasswordUserRepository {
	return &passwordUserRepoImpl{
		Db: c.Db,
	}
}
