package loader

import (
	"context"
	"github.com/graph-gophers/dataloader/v7"
	"github.com/mhmmdFsl/my-online-petshop/pet-product/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ctxKye string

const (
	loadersKey = ctxKye("dataloaders")
)

type shopReader struct {
	Collection *mongo.Collection
}

func (u *shopReader) getShop(ctx context.Context, productIds []string) []*dataloader.Result[*model.Shop] {
	
	result := make([]*dataloader.Result[*model.Shop], 0, len(productIds))
	
	return result
}
