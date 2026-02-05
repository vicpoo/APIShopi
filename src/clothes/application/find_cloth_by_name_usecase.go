// application/find_cloth_by_name_usecase.go
package application

import (
	repositories "github.com/vicpoo/apiShop/src/clothes/domain"
	"github.com/vicpoo/apiShop/src/clothes/domain/entities"
)

type FindClothByNameUseCase struct {
	repo repositories.IClothRepository
}

func NewFindClothByNameUseCase(repo repositories.IClothRepository) *FindClothByNameUseCase {
	return &FindClothByNameUseCase{repo: repo}
}

func (uc *FindClothByNameUseCase) Run(name string) ([]entities.Cloth, error) {
	return uc.repo.FindByName(name)
}