package service

import (
	"context"

	"github.com/dzikriqalampacil/cariilmu-course/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, UserId int64)
	FindById(ctx context.Context, UserId int64) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}
