package service

import (
	"context"
	"database/sql"

	"github.com/dzikriqalampacil/cariilmu-course/exception"
	"github.com/dzikriqalampacil/cariilmu-course/helper"
	"github.com/dzikriqalampacil/cariilmu-course/model/domain"
	"github.com/dzikriqalampacil/cariilmu-course/model/web"
	"github.com/dzikriqalampacil/cariilmu-course/repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(UserRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: UserRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	User := domain.User{
		Name: request.Name,
		Email: request.Email,
	}

	User = service.UserRepository.Save(ctx, tx, User)

	return helper.ToUserResponse(User)
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	User, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	User.Name = request.Name

	User = service.UserRepository.Update(ctx, tx, User)

	return helper.ToUserResponse(User)
}

func (service *UserServiceImpl) Delete(ctx context.Context, UserId int64) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	User, err := service.UserRepository.FindById(ctx, tx, UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRepository.Delete(ctx, tx, User)
}

func (service *UserServiceImpl) FindById(ctx context.Context, UserId int64) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	User, err := service.UserRepository.FindById(ctx, tx, UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(User)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.UserRepository.FindAll(ctx, tx)

	return helper.ToUserResponses(categories)
}
