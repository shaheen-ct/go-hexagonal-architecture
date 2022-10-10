package user

import "github.com/shaheen-ct/go-hexagonal-architecture/internal/entity"

type Repository interface {
	Create(user *entity.User) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Get(id int64) (*entity.User, error)
	Delete(id int64) error
}
