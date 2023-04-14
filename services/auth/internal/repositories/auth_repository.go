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
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
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
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505": // duplicate key error
				return nil, NewErrorWrapper(CodeClientError, "duplicate errror", fmt.Errorf("unique violation %w", err))
			default:

				return nil, NewErrorWrapper(CodeServerError, "pg error", fmt.Errorf("pg error %w", err))
			}

		}

	}

	return userData, nil
}

func (ar *authRepo) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user *models.User
	result := ar.db.Where("username = ?", username).First(&user)

	if result.RowsAffected <= 0 {
		if result.Error != nil {
			return nil, result.Error
		}
		return nil, gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
