package follow

import (
	"context"
	"twitter-uala/pkg/domain/follow/entities"
	"twitter-uala/pkg/repositories/follow/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) CreateFollow(ctx context.Context, follow entities.Follow) error {
	model := models.FromDomain(follow)
	_, err := r.database.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetFollowers(ctx context.Context, userID string) ([]*entities.Follow, error) {
	filter := bson.M{"followed_id": userID}
	cursor, err := r.database.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var results []*models.Follow
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}

	return models.ModelsToDomain(results), nil
}
