package repository

import (
	"context"

	"github.com/Andhika-GIT/Go-REST-Event-Management/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, tx *gorm.DB, user domain.User) error
	Update(ctx context.Context, tx *gorm.DB, user domain.User) error
	Delete(ctx context.Context, tx *gorm.DB, user domain.User) error
	FindAll(ctx context.Context, tx *gorm.DB, user *domain.User) error
	FindById(ctx context.Context, tx *gorm.DB, user domain.User) error
}
