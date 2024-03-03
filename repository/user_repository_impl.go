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

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, user domain.User) error {

	err := tx.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, user domain.User) error {

	err := tx.Save(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, user domain.User) error {

	err := tx.Where("id = ?", user.Id).Delete(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB, user *domain.User) error {

	var users []domain.User

	err := tx.Find(&users).Error

	user = users

	if err != nil {
		return err
	}

	return nil

}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, user domain.User) error {

	err := tx.First(&user, user.Id).Error

	if err != nil {
		return err
	}

	return nil

}
