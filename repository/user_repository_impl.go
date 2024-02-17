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

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error) {
	err := tx.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error) {
	err := tx.Where("id = ?", user.Id).Delete(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, userId int32) (domain.User, error) {

	user := domain.User{}
	err := tx.First(&user, userId).Error

	if err != nil {
		return user, err
	}

	return user, nil

}
