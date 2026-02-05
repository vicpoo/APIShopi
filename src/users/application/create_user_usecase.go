// application/create_user_usecase.go
package application

import (
	repositories "github.com/vicpoo/apiShop/src/users/domain"
	"github.com/vicpoo/apiShop/src/users/domain/entities"
)

type CreateUserUseCase struct {
	repo repositories.IUserRepository
}

func NewCreateUserUseCase(repo repositories.IUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}

func (uc *CreateUserUseCase) Run(user *entities.User) (*entities.User, error) {
	err := uc.repo.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}