package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repository) InsertOne(
	ctx context.Context,
	document interface{},
	opts ...*options.InsertOneOptions,
) (*mongo.InsertOneResult, error) {
	result, err := r.collection.InsertOne(ctx, document, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) Find(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOptions,
) (cur *mongo.Cursor, err error) {
	cusor, err := r.collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	return cusor, nil
}
