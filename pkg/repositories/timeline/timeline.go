package timeline

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"twitter-uala/pkg/domain/timeline/entities"
	"twitter-uala/pkg/repositories/timeline/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repository) CreateTimeline(ctx context.Context, timeline entities.Timeline) error {
	model := models.FromDomain(timeline)
	model.CreatedAt = time.Now()
	_, err := r.database.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CreateTimelineInCache(ctx context.Context, timeline entities.Timeline) error {
	previousTimeline, err := r.GetLatestsTimelineFromCache(ctx, timeline.UserID)
	if err != nil {
		return err
	}

	var timelineToSave models.TimelineCache
	if previousTimeline != nil {
		var ids []string
		for _, pt := range previousTimeline {
			ids = append(ids, pt.TweetID)
		}

		timelineToSave.TweetIDs = ids
	}

	timelineToSave.TweetIDs = append(timelineToSave.TweetIDs, timeline.TweetID)

	item, err := json.Marshal(timelineToSave)
	if err != nil {
		return err
	}

	r.cache.Set(fmt.Sprintf("timeline_%s", timeline.UserID), item)

	return nil
}

func (r *repository) GetLatestsTimeline(ctx context.Context, userID string) ([]*entities.Timeline, error) {
	filter := bson.M{"user_id": userID}
	opts := &options.FindOptions{
		Sort:  bson.M{"created_at": -1},
		Limit: &r.configs.LatestsTweetsLimit,
	}
	cursor, err := r.database.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var results []*models.Timeline
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}

	return models.ModelsToDomain(results), nil
}

func (r *repository) GetLatestsTimelineFromCache(ctx context.Context, userID string) ([]*entities.Timeline, error) {
	item := r.cache.Get(fmt.Sprintf("timeline_%s", userID))
	if item == nil || item.Value() == nil {
		return nil, nil
	}

	timelineRaw, ok := item.Value().([]byte)
	if !ok {
		return nil, errors.New("error parsing item from cache")
	}

	var timelineModel models.TimelineCache
	err := json.Unmarshal(timelineRaw, &timelineModel)
	if err != nil {
		return nil, fmt.Errorf("error parsing item from cache: %s", err.Error())
	}

	return timelineModel.ToDomain(), nil
}
