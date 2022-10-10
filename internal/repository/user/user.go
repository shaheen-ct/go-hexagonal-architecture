package user

import (
	"fmt"
	"sync"

	"github.com/shaheen-ct/go-hexagonal-architecture/internal/entity"

	"gorm.io/gorm"
)

var (
	repo     *Repository
	repoOnce sync.Once
)

type Repository struct {
	db *gorm.DB
}

func (r Repository) Create(user *entity.User) (*entity.User, error) {
	err := r.db.Debug().Create(user).Error
	return user, err
}

func (r Repository) Update(user *entity.User) (*entity.User, error) {
	fmt.Println("kiki", *user)

	err := r.db.Debug().Model(&user).Where("id=?", user.ID).Updates(*user).Error
	fmt.Println("kiko", err)
	return user, err
}

func (r Repository) Get(id int64) (*entity.User, error) {
	var user *entity.User
	err := r.db.Debug().Where("id =?", id).Find(&user).Error

	return user, err
}

func (r Repository) Delete(id int64) error {
	return r.db.Debug().Delete(&entity.User{}, id).Error
}

func ProvideRepository(db *gorm.DB) *Repository {
	repoOnce.Do(func() {
		repo = &Repository{
			db: db,
		}
	})
	return repo
}
