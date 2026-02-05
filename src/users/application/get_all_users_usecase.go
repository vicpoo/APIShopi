// application/get_all_users_usecase.go
package application

import (
	repositories "github.com/vicpoo/apiShop/src/users/domain"
	"github.com/vicpoo/apiShop/src/users/domain/entities"
)

type GetAllUsersUseCase struct {
	repo repositories.IUserRepository
}

func NewGetAllUsersUseCase(repo repositories.IUserRepository) *GetAllUsersUseCase {
	return &GetAllUsersUseCase{repo: repo}
}

func (uc *GetAllUsersUseCase) Run() ([]entities.User, error) {
	return uc.repo.GetAll()
}