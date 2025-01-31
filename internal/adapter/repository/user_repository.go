package repository

import (
	"context"

	"github.com/inyourtime/noti-me-server/internal/core/domain"
	"github.com/inyourtime/noti-me-server/internal/core/port"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user domain.User) (*domain.User, error) {
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindOne(ctx context.Context, condition interface{}) (*domain.User, error) {
	return nil, nil
}
