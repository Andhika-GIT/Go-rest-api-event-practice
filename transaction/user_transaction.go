package transaction

import (
	"context"

	"github.com/Andhika-GIT/Go-REST-Event-Management/model/domain"
	"github.com/Andhika-GIT/Go-REST-Event-Management/repository"
	"gorm.io/gorm"
)

// type UserModel interface {
// 	domain.User | []domain.User
// }

func RunTransactionType(transactionType string, ctx context.Context, tx *gorm.DB, UserRepository repository.UserRepository, user domain.User, optional []int32) error {
	var err error
	switch transactionType {
	case "create":
		err = UserRepository.Create(ctx, tx, user)
	case "update":
		err = UserRepository.Update(ctx, tx, user)
	case "delete":
		err = UserRepository.Delete(ctx, tx, user)
	case "findAll":
		err = UserRepository.FindAll(ctx, tx, &user)
	case "findById":
		for _, id := range optional {
			user.Id = id
			err = UserRepository.FindById(ctx, tx, user)
		}
	}

	return err

}

func UserDatabaseTransaction[T any](transactionType string, ctx context.Context, DB *gorm.DB, repository repository.UserRepository, optional ...int32) (T, error) {

	var user domain.User

	err := DB.Transaction(func(tx *gorm.DB) error {
		err := RunTransactionType(transactionType, ctx, tx, repository, user, optional)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return user, err
	}

	return user, nil
}
