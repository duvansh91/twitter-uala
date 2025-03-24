package mongodb

import (
	"context"
	"time"
	"twitter-uala/pkg/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository groups methods to interact with a MongoDB collection.
type Repository interface {
	InsertOne(
		ctx context.Context,
		document interface{},
		opts ...*options.InsertOneOptions,
	) (*mongo.InsertOneResult, error)
	Find(
		ctx context.Context,
		filter interface{},
		opts ...*options.FindOptions,
	) (cur *mongo.Cursor, err error)
}

type repository struct {
	collection *mongo.Collection
}

// NewRepository creates a new instance of MongoDB repository.
func NewRepository(
	ctx context.Context,
	conf *configs.MongoDBConfig,
	collection string,
) (Repository, error) {
	options := options.Client()
	options.ApplyURI(conf.Uri)
	options.SetConnectTimeout(time.Duration(conf.Timeout) * time.Second)

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &repository{
		collection: client.Database(conf.DatabaseName).Collection(collection),
	}, nil
}
