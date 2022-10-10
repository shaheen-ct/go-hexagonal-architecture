package user

import (
	"context"
	"sync"

	"github.com/shaheen-ct/go-hexagonal-architecture/internal/entity"
	"github.com/shaheen-ct/go-hexagonal-architecture/internal/interface/user"
)

var (
	srv     *Service
	srvOnce sync.Once
)

type Service struct {
	repo      user.Repository
	validator user.Validator
}

func (s Service) Create(_ context.Context, user *entity.User) (*entity.User, error) {
	return s.repo.Create(user)
}

func (s Service) Update(_ context.Context, user *entity.User) (*entity.User, error) {
	return s.repo.Update(user)

}

func (s Service) Get(_ context.Context, id int64) (*entity.User, error) {
	return s.repo.Get(id)
}

func (s Service) Delete(_ context.Context, id int64) error {
	return s.repo.Delete(id)
}

func ProvideService(validator user.Validator, repo user.Repository) *Service {
	srvOnce.Do(func() {
		srv = &Service{
			repo:      repo,
			validator: validator,
		}
	})

	return srv
}
