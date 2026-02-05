// application/create_cloth_usecase.go
package application

import (
	repositories "github.com/vicpoo/apiShop/src/clothes/domain"
	"github.com/vicpoo/apiShop/src/clothes/domain/entities"
)

type CreateClothUseCase struct {
	repo repositories.IClothRepository
}

func NewCreateClothUseCase(repo repositories.IClothRepository) *CreateClothUseCase {
	return &CreateClothUseCase{repo: repo}
}

func (uc *CreateClothUseCase) Run(cloth *entities.Cloth) (*entities.Cloth, error) {
	err := uc.repo.Save(cloth)
	if err != nil {
		return nil, err
	}
	return cloth, nil
}