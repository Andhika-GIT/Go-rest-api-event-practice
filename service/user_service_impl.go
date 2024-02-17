package service

import (
	"context"

	"github.com/Andhika-GIT/Go-REST-Event-Management/model/domain"
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

	user := domain.User{}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		_, err := service.UserRepository.Create(ctx, tx, user)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return web.UserResponse{}, err
	}

	return web.UserResponse{
		Id:   user.Id,
		Name: user.Name,
	}, nil

}
