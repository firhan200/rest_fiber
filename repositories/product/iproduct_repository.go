package repositories

import (
	"github.com/firhan200/rest_fiber/models"

	"gorm.io/gorm"
)

type IProductRepository interface {
	GetAll() ([]models.Product, error)
	Add(*models.Product) (bool, error)
	Update(id uint, name string, price int32) (bool, error)
	Delete(id uint) (bool, error)
}

func NewProductRepository(db *gorm.DB, impl IProductRepository) IProductRepository {
	return &ProductRepository{
		db: db,
	}
}
