// application/delete_cloth_usecase.go
package application

import repositories "github.com/vicpoo/apiShop/src/clothes/domain"

type DeleteClothUseCase struct {
	repo repositories.IClothRepository
}

func NewDeleteClothUseCase(repo repositories.IClothRepository) *DeleteClothUseCase {
	return &DeleteClothUseCase{repo: repo}
}

func (uc *DeleteClothUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}