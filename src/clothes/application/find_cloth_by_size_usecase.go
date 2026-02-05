// application/find_cloth_by_size_usecase.go
package application

import (
	repositories "github.com/vicpoo/apiShop/src/clothes/domain"
	"github.com/vicpoo/apiShop/src/clothes/domain/entities"
)

type FindClothBySizeUseCase struct {
	repo repositories.IClothRepository
}

func NewFindClothBySizeUseCase(repo repositories.IClothRepository) *FindClothBySizeUseCase {
	return &FindClothBySizeUseCase{repo: repo}
}

func (uc *FindClothBySizeUseCase) Run(size string) ([]entities.Cloth, error) {
	return uc.repo.FindBySize(size)
}