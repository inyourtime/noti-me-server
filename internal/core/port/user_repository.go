package port

import (
	"context"

	"github.com/inyourtime/noti-me-server/internal/core/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) (*domain.User, error)
	FindOne(ctx context.Context, cond interface{}) (*domain.User, error)
}
