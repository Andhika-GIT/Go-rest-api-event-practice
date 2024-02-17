package repository

import (
	"context"

	"github.com/Andhika-GIT/Go-REST-Event-Management/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error)
	Update(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error)
	Delete(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error)
	FindById(ctx context.Context, tx *gorm.DB, userId int32) (domain.User, error)
}
