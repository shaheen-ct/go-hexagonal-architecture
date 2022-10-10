package user

import (
	"context"

	"github.com/shaheen-ct/go-hexagonal-architecture/internal/entity"
)

type Service interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
	Get(ctx context.Context, id int64) (*entity.User, error)
	Delete(ctx context.Context, id int64) error
}
