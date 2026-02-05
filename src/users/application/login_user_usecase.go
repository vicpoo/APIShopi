// application/login_user_usecase.go
package application

import (
	repositories "github.com/vicpoo/apiShop/src/users/domain"
	"github.com/vicpoo/apiShop/src/users/domain/entities"
)

type LoginUserUseCase struct {
	repo repositories.IUserRepository
}

func NewLoginUserUseCase(repo repositories.IUserRepository) *LoginUserUseCase {
	return &LoginUserUseCase{repo: repo}
}

func (uc *LoginUserUseCase) Run(email, password string) (*entities.User, error) {
	return uc.repo.Login(email, password)
}