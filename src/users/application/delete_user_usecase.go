// application/delete_user_usecase.go
package application

import repositories "github.com/vicpoo/apiShop/src/users/domain"

type DeleteUserUseCase struct {
	repo repositories.IUserRepository
}

func NewDeleteUserUseCase(repo repositories.IUserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{repo: repo}
}

func (uc *DeleteUserUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}