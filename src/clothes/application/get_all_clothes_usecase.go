// application/get_all_clothes_usecase.go
package application

import (
	repositories "github.com/vicpoo/apiShop/src/clothes/domain"
	"github.com/vicpoo/apiShop/src/clothes/domain/entities"
)

type GetAllClothesUseCase struct {
	repo repositories.IClothRepository
}

func NewGetAllClothesUseCase(repo repositories.IClothRepository) *GetAllClothesUseCase {
	return &GetAllClothesUseCase{repo: repo}
}

func (uc *GetAllClothesUseCase) Run() ([]entities.Cloth, error) {
	return uc.repo.GetAll()
}