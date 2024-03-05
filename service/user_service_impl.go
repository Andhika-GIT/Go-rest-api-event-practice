package service

import (
	"context"

	"github.com/Andhika-GIT/Go-REST-Event-Management/model/domain"
	"github.com/Andhika-GIT/Go-REST-Event-Management/model/web"
	"github.com/Andhika-GIT/Go-REST-Event-Management/repository"
	"github.com/Andhika-GIT/Go-REST-Event-Management/transaction"
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

	user, err := transaction.UserDatabaseTransaction("create", ctx, service.DB, service.UserRepository, 0)

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

	user, err := transaction.UserDatabaseTransaction("update", ctx, service.DB, service.UserRepository, 0)

	if err != nil {
		return web.UserResponse{}, err
	}

	return web.UserResponse{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int32) error {
	_, err := transaction.UserDatabaseTransaction("delete", ctx, service.DB, service.UserRepository, 0)

	if err != nil {
		return err
	}

	return nil
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int32) (web.UserResponse, error) {
	user, err := transaction.UserDatabaseTransaction("findById", ctx, service.DB, service.UserRepository, userId)

	if err != nil {
		return web.UserResponse{}, nil
	}

	return web.UserResponse{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}

func (service *UserServiceImpl) FindAll(ctx context.Context) ([]web.UserResponse, error) {
	var users []domain.User

	var usersResponse []web.UserResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		err := service.UserRepository.FindAll(ctx, tx, users)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return usersResponse, nil
	}

	for _, value := range users {
		user := web.UserResponse{
			Id:   value.Id,
			Name: value.Name,
		}

		usersResponse = append(usersResponse, user)
	}

	return usersResponse, nil

}
