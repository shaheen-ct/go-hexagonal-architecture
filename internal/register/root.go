package register

import (
	"log"

	"github.com/shaheen-ct/go-hexagonal-architecture/db/pgx"
	"github.com/shaheen-ct/go-hexagonal-architecture/internal/interface/user"
	"github.com/shaheen-ct/go-hexagonal-architecture/internal/wirec"
)

type RegistryOptions struct {
	PgxOption pgx.Option
}

type Registry struct {
	User user.Controller
}

type RegistryOptionFunc func() RegistryOptions

func (fn RegistryOptionFunc) Must() RegistryOptions {
	opt := fn()
	if opt.PgxOption == (pgx.Option{}) {
		log.Fatalln("missing postgres database options")
	}
	return opt
}

func Register(fn RegistryOptionFunc) Registry {
	// validate RegistryOptions
	opt := fn.Must()

	// DB setup
	db := pgx.Setup(opt.PgxOption)()
	return Registry{
		User: wirec.UserWire(db),
	}
}
