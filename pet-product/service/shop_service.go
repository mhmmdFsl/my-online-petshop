package service

import (
	"context"
	"github.com/dgryski/trifles/uuid"
	"github.com/mhmmdFsl/my-online-petshop/pet-product/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ShopService interface {
	CreateShop(i *model.NewShop) (*model.Shop, error)
	GetAll(f *model.ShopFilter) ([]*model.Shop, error)
}

type shopServiceImpl struct {
	Collection *mongo.Collection
}

func (s shopServiceImpl) GetAll(f *model.ShopFilter) ([]*model.Shop, error) {
	findOptions := options.Find()
	limit := int64(f.Limit)
	findOptions.Limit = &limit
	filter := bson.D{}
	var regex primitive.Regex
	if f.Name != nil {
		regex = primitive.Regex{Pattern: *f.Name, Options: "i"}
		filter = append(filter, bson.E{Key: "name", Value: regex})
	}

	if f.UserID != nil {
		regex = primitive.Regex{Pattern: *f.UserID, Options: "i"}
		filter = append(filter, bson.E{Key: "userid", Value: regex})
	}

	var rs []*model.Shop
	cur, err := s.Collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var p model.Shop
		err := cur.Decode(&p)
		if err != nil {
			return nil, err
		}
		rs = append(rs, &p)
	}

	if cur.Err() != nil {
		return nil, err
	}

	err = cur.Close(context.TODO())
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (s shopServiceImpl) CreateShop(i *model.NewShop) (*model.Shop, error) {

	shop := &model.Shop{
		ID:         uuid.UUIDv4(),
		Name:       i.Name,
		UserID:     i.UserID,
		LogoURL:    i.LogoURL,
		IsVerified: false,
		ProductID:  nil,
		Status:     "ACTIVE",
		CreatedAt:  time.Now().Format(time.RFC3339Nano),
		UpdatedAt:  time.Now().Format(time.RFC3339Nano),
	}

	_, err := s.Collection.InsertOne(context.TODO(), shop)
	if err != nil {
		return nil, err
	}

	return shop, nil
}

type ShopServiceCfg struct {
	Collection *mongo.Collection
}

func NewShopService(cfg *ShopServiceCfg) ShopService {
	return &shopServiceImpl{
		Collection: cfg.Collection,
	}
}
