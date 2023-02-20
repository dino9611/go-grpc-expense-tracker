package repositories

import (
	"context"
	"errors"
	"fmt"
	"grpc-finance-app/services/auth/internal/models"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type IAuthRepo interface {
	AddUser(ctx context.Context, userData *models.User) (*models.User, error)
}

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) IAuthRepo {
	return &authRepo{
		db: db,
	}
}

func (ar *authRepo) AddUser(ctx context.Context, userData *models.User) (*models.User, error) {

	if err := ar.db.WithContext(ctx).Create(&userData).Error; err != nil {

		var pgErr *pgconn.PgError

		// fmt.Println("tes", err)
		fmt.Println("tes", err)
		if errors.As(err, &pgErr) {
			fmt.Println("tes2", err)
			switch pgErr.Code {
			case "23505": // duplicate key error
				fmt.Println("tes2", err)
				return nil, NewErrorWrapper(CodeClientError, "duplicate errror", fmt.Errorf("unique violation %w", err))
			default:

				return nil, NewErrorWrapper(CodeServerError, "pg error", fmt.Errorf("pg error %w", err))
			}

		}
		// pgErr := err.(*pgconn.PgError)
		// if errors.Is(err, pgErr) {
		// 	fmt.Println("tes2", err)
		// 	switch pgErr.Code {
		// 	case "23505": // duplicate key error
		// 		fmt.Println("tes2", err)
		// 		return nil, fmt.Errorf("unique violation %w", err)
		// 	default:

		// 		return nil, fmt.Errorf("pg error %w", err)
		// 	}

		// }
	}

	return userData, nil

}
