//go:build wireinject
// +build wireinject

package wirec

import (
	c "github.com/shaheen-ct/go-hexagonal-architecture/api/graph/controller/user"
	i "github.com/shaheen-ct/go-hexagonal-architecture/internal/interface/user"
	r "github.com/shaheen-ct/go-hexagonal-architecture/internal/repository/user"
	s "github.com/shaheen-ct/go-hexagonal-architecture/internal/service/user"

	"gorm.io/gorm"

	"github.com/google/wire"
)

var userProviderSet = wire.NewSet(
	r.ProvideRepository,
	s.ProvideValidator,
	s.ProvideService,
	c.ProvideController,
	wire.Struct(new(c.DTO), "*"),
	//// bind each one of the interfaces
	wire.Bind(new(i.Service), new(*s.Service)),
	wire.Bind(new(i.Validator), new(*s.Validator)),
	wire.Bind(new(i.Repository), new(*r.Repository)),
)

func UserWire(db *gorm.DB) *c.Controller {
	panic(wire.Build(
		userProviderSet,
	))
}
