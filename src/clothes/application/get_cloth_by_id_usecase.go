// application/get_cloth_by_id_usecase.go
package application

import (
	repositories "github.com/vicpoo/apiShop/src/clothes/domain"
	"github.com/vicpoo/apiShop/src/clothes/domain/entities"
)

type GetClothByIDUseCase struct {
	repo repositories.IClothRepository
}

func NewGetClothByIDUseCase(repo repositories.IClothRepository) *GetClothByIDUseCase {
	return &GetClothByIDUseCase{repo: repo}
}

func (uc *GetClothByIDUseCase) Run(id int32) (*entities.Cloth, error) {
	return uc.repo.GetByID(id)
}