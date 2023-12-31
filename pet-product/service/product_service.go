package service

import (
	"context"
	"github.com/dgryski/trifles/uuid"
	"github.com/mhmmdFsl/my-online-petshop/pet-product/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ProductService interface {
	Create(m *model.NewProduct) (*model.Product, error)
	GetAll() ([]*model.Product, error)
	Delete(id string) (string, error)
	UpdateProduct(up *model.UpdateProduct) (*model.Product, error)
}

type productServiceImpl struct {
	Collection *mongo.Collection
}

func (p productServiceImpl) UpdateProduct(up *model.UpdateProduct) (*model.Product, error) {
	filter := bson.D{
		{"id", up.ID},
	}
	var product *model.Product
	err := p.Collection.FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		return nil, err
	}
	
	if up.Name != "" {
		product.Name = up.Name
	} 

	if up.Price != nil {
		product.Price = *up.Price
	} 

	if up.ImageURL != nil {
		product.ImageURL = *up.ImageURL
	} 
	update := bson.M{
		"$set": bson.M{
			"name": product.Name,
			"price": product.Price,
			"imageUrl": product.ImageURL,
		},
	}

	_, err = p.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err	
	}
	return product, nil
}

func (p productServiceImpl) Delete(id string) (string, error) {
	_, err := p.Collection.DeleteOne(context.TODO(), bson.D{{"id", id}})
	if err != nil {
		return "", err
	}
	return "Success delete product", nil
}

func (p productServiceImpl) GetAll() ([]*model.Product, error) {
	findOptions := options.Find()
	findOptions.SetLimit(20)
	var rs []*model.Product
	cur, err := p.Collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var p model.Product
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

func (p productServiceImpl) Create(m *model.NewProduct) (*model.Product, error) {
	product := model.Product{
		ID:        uuid.UUIDv4(),
		Name:      m.Name,
		Price:     m.Price,
		ImageURL:  m.ImageURL,
		CreatedAt: time.Now().Format(time.RFC3339Nano),
		UpdatedAt: time.Now().Format(time.RFC3339Nano),
	}
	
	_, err := p.Collection.InsertOne(context.TODO(), product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

type ProductServiceCfg struct {
	Collection *mongo.Collection
}

func NewProductService(c *ProductServiceCfg) ProductService {
	return &productServiceImpl{
		Collection: c.Collection,
	}
}