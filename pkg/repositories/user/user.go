package user

import (
	"context"
	"errors"
	"twitter-uala/pkg/constants"
	"twitter-uala/pkg/domain/user/entities"
	"twitter-uala/pkg/repositories/user/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repository) CreateUser(ctx context.Context, user entities.User) (*entities.User, error) {
	model := models.FromDomain(user)
	_, err := r.database.InsertOne(ctx, model)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) FindUserByID(ctx context.Context, userID string) (*entities.User, error) {
	filter := bson.M{"user_id": userID}
	singleResult, err := r.database.FindOne(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New(constants.UserNotFoundError)
		}
		return nil, err
	}

	var model *models.User
	singleResult.Decode(&model)
	if err != nil {
		return nil, err
	}

	return model.ToDomain(), nil
}
