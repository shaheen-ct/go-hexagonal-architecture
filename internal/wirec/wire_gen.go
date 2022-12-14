// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wirec

import (
	"github.com/shaheen-ct/go-hexagonal-architecture/api/graph/controller/user"
	user4 "github.com/shaheen-ct/go-hexagonal-architecture/internal/interface/user"
	user2 "github.com/shaheen-ct/go-hexagonal-architecture/internal/repository/user"
	user3 "github.com/shaheen-ct/go-hexagonal-architecture/internal/service/user"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// Injectors from user.go:

func UserWire(db *gorm.DB) *user.Controller {
	repository := user2.ProvideRepository(db)
	validator := user3.ProvideValidator(repository)
	service := user3.ProvideService(validator, repository)
	dto := user.DTO{}
	controller := user.ProvideController(service, dto)
	return controller
}

// user.go:

var userProviderSet = wire.NewSet(user2.ProvideRepository, user3.ProvideValidator, user3.ProvideService, user.ProvideController, wire.Struct(new(user.DTO), "*"), wire.Bind(new(user4.Service), new(*user3.Service)), wire.Bind(new(user4.Validator), new(*user3.Validator)), wire.Bind(new(user4.Repository), new(*user2.Repository)))
