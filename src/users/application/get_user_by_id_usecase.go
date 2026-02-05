// application/get_user_by_id_usecase.go
package application

import (
	repositories "github.com/vicpoo/apiShop/src/users/domain"
	"github.com/vicpoo/apiShop/src/users/domain/entities"
)

type GetUserByIDUseCase struct {
	repo repositories.IUserRepository
}

func NewGetUserByIDUseCase(repo repositories.IUserRepository) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{repo: repo}
}

func (uc *GetUserByIDUseCase) Run(id int32) (*entities.User, error) {
	return uc.repo.GetByID(id)
}