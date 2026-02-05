package infrastructure

import (
	"github.com/vicpoo/apiShop/src/clothes/application"
)

func InitClothDependencies() (
	*CreateClothController,
	*UpdateClothController,
	*DeleteClothController,
	*GetClothByIDController,
	*GetAllClothesController,
	*FindClothByNameController,
	*FindClothBySizeController,
	*FindClothByPriceRangeController,
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLClothRepository()

	// Casos de uso
	createUseCase := application.NewCreateClothUseCase(repo)
	updateUseCase := application.NewUpdateClothUseCase(repo)
	deleteUseCase := application.NewDeleteClothUseCase(repo)
	getByIDUseCase := application.NewGetClothByIDUseCase(repo)
	getAllUseCase := application.NewGetAllClothesUseCase(repo)
	findByNameUseCase := application.NewFindClothByNameUseCase(repo)
	findBySizeUseCase := application.NewFindClothBySizeUseCase(repo)
	findByPriceRangeUseCase := application.NewFindClothByPriceRangeUseCase(repo)

	// Controladores (ahora manejan archivos)
	createController := NewCreateClothController(createUseCase)
	updateController := NewUpdateClothController(updateUseCase)
	deleteController := NewDeleteClothController(deleteUseCase)
	getByIDController := NewGetClothByIDController(getByIDUseCase)
	getAllController := NewGetAllClothesController(getAllUseCase)
	findByNameController := NewFindClothByNameController(findByNameUseCase)
	findBySizeController := NewFindClothBySizeController(findBySizeUseCase)
	findByPriceRangeController := NewFindClothByPriceRangeController(findByPriceRangeUseCase)

	return createController, updateController, deleteController,
		getByIDController, getAllController, findByNameController,
		findBySizeController, findByPriceRangeController
}