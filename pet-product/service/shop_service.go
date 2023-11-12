package service

import (
	"github.com/mhmmdFsl/my-online-petshop/pet-product/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShopService interface {
	CreateShop(i *model.NewShop) (*model.Shop, error)
}

type shopServiceImpl struct {
	Collection mongo.Collection
}

func (s shopServiceImpl) CreateShop(i *model.NewShop) (*model.Shop, error) {
	//TODO implement me
	panic("implement me")
}

type ShopServiceCfg struct {
	Collection mongo.Collection
}

func NewShopService(cfg *ShopServiceCfg) ShopService {
	return &shopServiceImpl{
		Collection: cfg.Collection,
	}
}
