package models

import "twitter-uala/pkg/domain/user/entities"

type User struct {
	UserID string `bson:"user_id"`
}

func (u *User) ToDomain() *entities.User {
	return &entities.User{
		UserID: u.UserID,
	}
}

func FromDomain(user entities.User) User {
	return User{UserID: user.UserID}
}
