package user

import (
	"context"
	"twitter-uala/pkg/domain/user/entities"
)

func (r *repository) CreateUser(ctx context.Context, user entities.User) error {
	return nil
}

func (r *repository) FindUser(ctx context.Context, userName string) error {
	return nil
}
