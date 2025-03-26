package dto

type TweetResponse struct {
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}

type PublishTweetRequest struct {
	Content string `json:"content" validate:"required"`
}
