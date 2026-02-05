// application/update_user_usecase.go
package application

import (
	repositories "github.com/vicpoo/apiShop/src/users/domain"
	"github.com/vicpoo/apiShop/src/users/domain/entities"
)

type UpdateUserUseCase struct {
	repo repositories.IUserRepository
}

func NewUpdateUserUseCase(repo repositories.IUserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{repo: repo}
}

func (uc *UpdateUserUseCase) Run(user *entities.User) (*entities.User, error) {
	err := uc.repo.Update(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}