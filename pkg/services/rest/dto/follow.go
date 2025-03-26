package dto

type FollowRequest struct {
	UserToFollow string `json:"user_id" validation:"required"`
}
