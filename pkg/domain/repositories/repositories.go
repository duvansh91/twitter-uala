package repositories

import (
	"context"
	"twitter-uala/pkg/domain/user/entities"
)

type (
	User interface {
		CreateUser(ctx context.Context, user entities.User) error
		FindUser(ctx context.Context, userName string) error
	}
)
