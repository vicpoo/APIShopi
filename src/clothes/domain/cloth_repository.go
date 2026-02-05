// cloth_repository.go
package domain

import (
	"github.com/vicpoo/apiShop/src/clothes/domain/entities"
)

type IClothRepository interface {
	// CRUD básico
	Save(cloth *entities.Cloth) error
	Update(cloth *entities.Cloth) error
	Delete(id int32) error
	GetByID(id int32) (*entities.Cloth, error)
	GetAll() ([]entities.Cloth, error)
	
	// Métodos de búsqueda
	FindByName(name string) ([]entities.Cloth, error)
	FindBySize(size string) ([]entities.Cloth, error)
	FindByPriceRange(minPrice, maxPrice float64) ([]entities.Cloth, error)
	
}