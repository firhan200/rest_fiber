package repositories

import (
	"errors"

	"github.com/firhan200/rest_fiber/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func (pr *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product

	//get from database
	pr.db.Model(&models.Product{}).Order("id desc").Find(&products)

	return products, nil
}

func (pr *ProductRepository) Add(p *models.Product) (bool, error) {
	// Create
	result := pr.db.Create(&p)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (pr *ProductRepository) Update(id uint, name string, price int32) (bool, error) {
	//get data
	var product models.Product
	result := pr.db.First(&product, id)

	if result.RowsAffected < 1 {
		return false, errors.New("Product not found")
	}

	if result.Error != nil {
		return false, result.Error
	}

	product.Name = name
	product.Price = price
	pr.db.Save(product)

	return true, nil
}

func (pr *ProductRepository) Delete(id uint) (bool, error) {
	//get data
	var product models.Product
	result := pr.db.First(&product, id)

	if result.RowsAffected < 1 {
		return false, errors.New("Product not found")
	}

	if result.Error != nil {
		return false, result.Error
	}

	pr.db.Delete(&product)

	return true, nil
}
