package store

import (
	"context"
	"fmt"

	"github.com/zeindevs/gocommerce/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductStore struct {
	db   *mongo.Database
	coll string
}

func NewMongoProductStore(db *mongo.Database) ProductStorer {

	return &MongoProductStore{
		db:   db,
		coll: "products",
	}
}

// GetByID implements ProductStorer.
func (m *MongoProductStore) GetByID(ctx context.Context, id string) (*types.Product, error) {
	p := m.db.Collection(m.coll).FindOne(ctx, bson.M{id: id})
	if p.Err() != nil {
		return nil, p.Err()
	}

	var pd types.Product
	if err := p.Decode(&pd); err != nil {
		return nil, fmt.Errorf("could not parse product")
	}

	return &pd, nil
}

// Insert implements ProductStorer.
func (m *MongoProductStore) Insert(ctx context.Context, p *types.Product) error {
	res, err := m.db.Collection(m.coll).InsertOne(ctx, p)
	if err != nil {
		return err
	}

	p.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return err
}

// GetAll implements ProductStorer.
func (m *MongoProductStore) GetAll(ctx context.Context) ([]*types.Product, error) {
	res, err := m.db.Collection(m.coll).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var pd []*types.Product
	err = res.All(ctx, pd)

	return pd, err
}

// Update implements ProductStorer.
func (m *MongoProductStore) Update(ctx context.Context, p *types.Product) error {
	panic("unimplemented")
}

// Delete implements ProductStorer.
func (m *MongoProductStore) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}
