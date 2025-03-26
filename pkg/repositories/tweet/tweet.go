package tweet

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"twitter-uala/pkg/domain/tweet/entities"
	"twitter-uala/pkg/repositories/tweet/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) CreateTweet(ctx context.Context, tweet entities.Tweet) (*entities.Tweet, error) {
	model := models.FromDomain(tweet)
	model.CreatedAt = time.Now()
	_, err := r.database.InsertOne(ctx, model)
	if err != nil {
		return nil, err
	}

	createdTweet := &entities.Tweet{
		TweetID: model.TweetID,
		UserID:  model.UserID,
		Content: model.Content,
	}

	return createdTweet, nil
}

func (r *repository) CreateTweetInCache(ctx context.Context, tweet entities.Tweet) error {
	tweetToSave := models.FromDomain(tweet)

	item, err := json.Marshal(tweetToSave)
	if err != nil {
		return err
	}

	r.cache.Set(fmt.Sprintf("tweet_%s", tweetToSave.TweetID), item)

	return nil
}

func (r *repository) GetTweetsInBatch(ctx context.Context, tweetIDs []string) ([]*entities.Tweet, error) {
	filter := bson.M{"tweet_id": bson.M{"$in": tweetIDs}}
	cursor, err := r.database.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var results []*models.Tweet
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}

	return models.ModelsToDomain(results), nil
}

func (r *repository) GetTweetsInBatchFromCache(ctx context.Context, tweetIDs []string) ([]*entities.Tweet, error) {
	var tweetEntities []*entities.Tweet
	for _, id := range tweetIDs {
		item := r.cache.Get(fmt.Sprintf("tweet_%s", id))
		if item == nil || item.Value() == nil {
			return nil, nil
		}

		tweetRaw, ok := item.Value().([]byte)
		if !ok {
			return nil, errors.New("error parsing item from cache")
		}

		var tweetModel models.Tweet
		err := json.Unmarshal(tweetRaw, &tweetModel)
		if err != nil {
			return nil, fmt.Errorf("error parsing item from cache: %s", err.Error())
		}

		tweetEntities = append(tweetEntities, tweetModel.ToDomain())
	}

	return tweetEntities, nil
}
