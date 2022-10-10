package user

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/shaheen-ct/go-hexagonal-architecture/api/graph/models"
	"github.com/shaheen-ct/go-hexagonal-architecture/internal/interface/user"
)

var (
	ctr     *Controller
	ctrOnce sync.Once
)

// Controller responsible for DTO and logging there error in the system
type Controller struct {
	srv user.Service
	dto DTO
}

func (c Controller) Create(ctx context.Context, user models.UserInput) (*models.User, error) {
	dto, err := c.dto.toUserEntity(user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rs, err := c.srv.Create(ctx, dto)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return c.dto.toUserModel(rs), err
}

func (c Controller) Update(ctx context.Context, user models.UserInput) (*models.User, error) {
	dto, err := c.dto.toUserEntity(user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rs, err := c.srv.Update(ctx, dto)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return c.dto.toUserModel(rs), err
}

func (c Controller) Get(ctx context.Context, id string) (*models.User, error) {
	i, err := strconv.ParseInt(id, 0, 10)
	if err != nil {
		return nil, err
	}
	rs, err := c.srv.Get(ctx, i)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return c.dto.toUserModel(rs), err
}

func (c Controller) Delete(ctx context.Context, id string) error {
	i, err := strconv.ParseInt(id, 0, 10)
	if err != nil {
		return err
	}
	err = c.srv.Delete(ctx, i)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func ProvideController(srv user.Service, dto DTO) *Controller {
	ctrOnce.Do(func() {
		ctr = &Controller{
			srv: srv,
			dto: dto,
		}
	})
	return ctr
}
