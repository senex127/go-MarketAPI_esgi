package product

import (
	"gorm.io/gorm"
)	

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (pr *ProductRepository) GetAll() []Product {
	var products []Product
	pr.db.Find(&products)
	return products
}

func (pr *ProductRepository) GetById(id int) Product {
	var product Product
	pr.db.First(&product, id)
	return product
}

func (pr *ProductRepository) Create(product Product) Product {
	pr.db.Create(&product)
	return product
}

func (pr *ProductRepository) Update(product Product) Product {
	pr.db.Save(&product)
	return product
}