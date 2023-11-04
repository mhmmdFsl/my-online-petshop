package repository

import (
	"auth-server/model"
	"gorm.io/gorm"
)

type PrincipalUserRepository interface {
	Save(rq *model.PrincipalUser) error
	FindByTypeAndValue(pt string, pv string) (*model.PrincipalUser, error)
	ExistsByValue(pv string) (bool, error)
}

type principalUserRepository struct {
	Db *gorm.DB
}

func (p principalUserRepository) ExistsByValue(pv string) (bool, error) {
	var count int64
	rs := p.Db.Where("principal_value = ?", pv).Find(&model.PrincipalUser{}).Count(&count)
	if err := rs.Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (p principalUserRepository) FindByTypeAndValue(pt string, pv string) (*model.PrincipalUser, error) {
	var pu model.PrincipalUser
	rs := p.Db.Where(&model.PrincipalUser{
		PrincipalType:  pt,
		PrincipalValue: pv,
	}).First(&pu)

	if err := rs.Error; err != nil {
		return nil, err
	}

	return &pu, nil
}

func (p principalUserRepository) Save(rq *model.PrincipalUser) error {
	rs := p.Db.Save(rq)
	if err := rs.Error; err != nil {
		return err
	}
	return nil
}

type PrincipalUserRepositoryCfg struct {
	Db *gorm.DB
}

func NewPrincipalUserRepository(c *PrincipalUserRepositoryCfg) PrincipalUserRepository {
	return &principalUserRepository{
		Db: c.Db,
	}
}
