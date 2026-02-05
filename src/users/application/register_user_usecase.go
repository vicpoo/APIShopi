// application/register_user_usecase.go
package application

import (
	repositories "github.com/vicpoo/apiShop/src/users/domain"
	"github.com/vicpoo/apiShop/src/users/domain/entities"
)

type RegisterUserUseCase struct {
	repo repositories.IUserRepository
}

func NewRegisterUserUseCase(repo repositories.IUserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{repo: repo}
}

func (uc *RegisterUserUseCase) Run(user *entities.User) (*entities.User, error) {
	err := uc.repo.Register(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}