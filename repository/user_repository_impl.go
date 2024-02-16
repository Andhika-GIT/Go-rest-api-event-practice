package repository

import (
	"context"

	"github.com/Andhika-GIT/Go-REST-Event-Management/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error) {
	err := tx.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
