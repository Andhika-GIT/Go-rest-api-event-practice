package service

import (
	"context"

	"github.com/Andhika-GIT/Go-REST-Event-Management/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) (web.UserResponse, error)
	Update(ctx context.Context, request web.UserUpdateRequest) (web.UserResponse, error)
	Delete(ctx context.Context, userId int32) error
	FindById(ctx context.Context, userId int32) (web.UserResponse, error)
}
