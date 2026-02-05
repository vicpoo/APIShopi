// application/update_cloth_usecase.go
package application

import (
	repositories "github.com/vicpoo/apiShop/src/clothes/domain"
	"github.com/vicpoo/apiShop/src/clothes/domain/entities"
)

type UpdateClothUseCase struct {
	repo repositories.IClothRepository
}

func NewUpdateClothUseCase(repo repositories.IClothRepository) *UpdateClothUseCase {
	return &UpdateClothUseCase{repo: repo}
}

func (uc *UpdateClothUseCase) Run(cloth *entities.Cloth) (*entities.Cloth, error) {
	err := uc.repo.Update(cloth)
	if err != nil {
		return nil, err
	}
	return cloth, nil
}