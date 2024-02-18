package service

import (
	"context"

	"github.com/Andhika-GIT/Go-REST-Event-Management/helper"
	"github.com/Andhika-GIT/Go-REST-Event-Management/model/web"
	"github.com/Andhika-GIT/Go-REST-Event-Management/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	validate       *validator.Validate
}

func NewUserService(UserRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: UserRepository,
		DB:             DB,
		validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) (web.UserResponse, error) {

	err := service.validate.Struct(request)

	if err != nil {

		return web.UserResponse{}, err
	}

	user, err := helper.DatabaseTransaction("create", ctx, service, 0)

	if err != nil {
		return web.UserResponse{}, err
	}

	return web.UserResponse{
		Id:   user.Id,
		Name: user.Name,
	}, nil

}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) (web.UserResponse, error) {
	err := service.validate.Struct(request)

	if err != nil {
		return web.UserResponse{}, err
	}

	user, err := helper.DatabaseTransaction("update", ctx, service, 0)

	if err != nil {
		return web.UserResponse{}, err
	}

	return web.UserResponse{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int32) error {
	_, err := helper.DatabaseTransaction("delete", ctx, service, 0)

	if err != nil {
		return err
	}

	return nil
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int32) (web.UserResponse, error) {
	user, err := helper.DatabaseTransaction("findById", ctx, service, userId)

	if err != nil {
		return web.UserResponse{}, nil
	}

	return web.UserResponse{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}
