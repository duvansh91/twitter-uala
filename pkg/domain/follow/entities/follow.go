package entities

type Follow struct {
	FollowerID string `bson:"follower_id"`
	FollowedID string `bson:"followed_id"`
}
