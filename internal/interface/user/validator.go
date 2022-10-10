package user

import "github.com/shaheen-ct/go-hexagonal-architecture/internal/entity"

type Validator interface {
	Valid(user *entity.User) error
}
