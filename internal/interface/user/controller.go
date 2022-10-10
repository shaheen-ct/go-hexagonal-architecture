package user

import (
	"context"

	"github.com/shaheen-ct/go-hexagonal-architecture/api/graph/models"
)

type Controller interface {
	Create(ctx context.Context, user models.UserInput) (*models.User, error)
	Update(ctx context.Context, user models.UserInput) (*models.User, error)
	Get(ctx context.Context, id string) (*models.User, error)
	Delete(ctx context.Context, id string) error
}
