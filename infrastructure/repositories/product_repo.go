package repositories

import (
	"ApiHEXaGONAL/config"
	"ApiHEXaGONAL/domain"
)

func GetProducts() []domain.Product {
	var products []domain.Product
	config.DB.Find(&products)
	return products
}

func CreateProduct(product *domain.Product) error {
	result := config.DB.Create(product)
	return result.Error
}

func GetProductByID(id uint) (domain.Product, error) {
	var product domain.Product
	result := config.DB.First(&product, id)
	return product, result.Error
}

func UpdateProduct(product *domain.Product) error {
	result := config.DB.Save(product)
	return result.Error
}

func DeleteProduct(id uint) error {
	result := config.DB.Delete(&domain.Product{}, id)
	return result.Error
}

type ProductRepository struct {
	// Puedes agregar campos aquí si es necesario
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

// CreateProduct crea un nuevo producto
func (r *ProductRepository) CreateProduct(product *domain.Product) error {
	result := config.DB.Create(product)
	return result.Error
}

func (r *ProductRepository) GetProductByID(id uint) (*domain.Product, error) {
	var product domain.Product
	result := config.DB.Preload("User").First(&product, id) // Preload carga el usuario
	return &product, result.Error
}
