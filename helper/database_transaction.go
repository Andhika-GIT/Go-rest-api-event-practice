package helper

import (
	"context"

	"github.com/Andhika-GIT/Go-REST-Event-Management/model/domain"
	"github.com/Andhika-GIT/Go-REST-Event-Management/repository"
	"github.com/Andhika-GIT/Go-REST-Event-Management/service"
	"gorm.io/gorm"
)

func RunTransactionType(transactionType string, ctx context.Context, tx *gorm.DB, UserRepository repository.UserRepository, user domain.User, optional []int32) error {
	var err error
	switch transactionType {
	case "create":
		_, err = UserRepository.Create(ctx, tx, user)
	case "update":
		_, err = UserRepository.Update(ctx, tx, user)
	case "delete":
		_, err = UserRepository.Delete(ctx, tx, user)
	case "findById":
		for _, id := range optional {
			_, err = UserRepository.FindById(ctx, tx, id)
		}
	}

	return err

}

func DatabaseTransaction(transactionType string, ctx context.Context, service service.UserServiceImpl, optional ...int32) (domain.User, error) {
	user := domain.User{}

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		err := RunTransactionType(transactionType, ctx, tx, service.UserRepository, user, optional)

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
