package user

import (
	"sync"

	"github.com/shaheen-ct/go-hexagonal-architecture/internal/entity"
	"github.com/shaheen-ct/go-hexagonal-architecture/internal/interface/user"
)

var (
	val     *Validator
	valOnce sync.Once
)

// Validator for business validation
type Validator struct {
	repo user.Repository
}

func (v Validator) Valid(user *entity.User) error {
	return nil
}

func ProvideValidator(repo user.Repository) *Validator {
	valOnce.Do(func() {
		val = &Validator{
			repo: repo,
		}
	})

	return val
}
