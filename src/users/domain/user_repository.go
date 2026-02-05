// user_repository.go
package domain

import (
	"github.com/vicpoo/apiShop/src/users/domain/entities"
)

type IUserRepository interface {

	Save(user *entities.User) error
	Update(user *entities.User) error
	Delete(id int32) error
	GetByID(id int32) (*entities.User, error)
	GetAll() ([]entities.User, error)
	
	Register(user *entities.User) error
	Login(email, password string) (*entities.User, error)
	
}