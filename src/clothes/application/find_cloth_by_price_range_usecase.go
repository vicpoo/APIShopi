// application/find_cloth_by_price_range_usecase.go
package application

import (
	repositories "github.com/vicpoo/apiShop/src/clothes/domain"
	"github.com/vicpoo/apiShop/src/clothes/domain/entities"
)

type FindClothByPriceRangeUseCase struct {
	repo repositories.IClothRepository
}

func NewFindClothByPriceRangeUseCase(repo repositories.IClothRepository) *FindClothByPriceRangeUseCase {
	return &FindClothByPriceRangeUseCase{repo: repo}
}

func (uc *FindClothByPriceRangeUseCase) Run(minPrice, maxPrice float64) ([]entities.Cloth, error) {
	return uc.repo.FindByPriceRange(minPrice, maxPrice)
}