package writer

import "twitter-uala/pkg/domain/user"

type Handler struct {
	userService user.Service
}

func NewHandler(userServ user.Service) *Handler {
	return &Handler{
		userService: userServ,
	}
}
