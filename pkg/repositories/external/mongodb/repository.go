package mongodb

import (
	"context"
	"fmt"
	"os"
	"time"
	"twitter-uala/pkg/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	FindOne(
		ctx context.Context,
		filter interface{},
		opts ...*options.FindOneOptions,
	) (cur *mongo.SingleResult, err error)
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(
	conf *configs.MongoDBConfig,
	collection string,
) Repository {
	ctx := context.TODO()
	options := options.Client()
	options.ApplyURI(os.Getenv(conf.Uri))
	options.SetConnectTimeout(time.Duration(conf.Timeout) * time.Second)

	if collection == "" {
		panic("mongodb collection is empty")
	}

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		panic(fmt.Sprintf("error connecting to db: %s", err.Error()))
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(fmt.Sprintf("error on ping to db: %s", err.Error()))
	}

	return &repository{
		collection: client.Database(os.Getenv(conf.DatabaseName)).Collection(collection),
	}
}
