package application

import (
	"ApiHEXaGONAL/config"
	"ApiHEXaGONAL/domain"
	"ApiHEXaGONAL/infrastructure/repositories"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService(productRepo *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

// GetProducts obtiene todos los productos
func (s *ProductService) GetProducts() ([]domain.Product, error) {
	var products []domain.Product
	result := config.DB.Find(&products)
	return products, result.Error
}

// CreateProduct crea un nuevo producto
func (s *ProductService) CreateProduct(product *domain.Product) error {
	result := config.DB.Create(product)
	return result.Error
}

// GetProductByID obtiene un producto por ID
func (s *ProductService) GetProductByID(id uint) (*domain.Product, error) {
	var product domain.Product
	result := config.DB.First(&product, id)
	return &product, result.Error
}

// UpdateProduct actualiza un producto existente
func (s *ProductService) UpdateProduct(product *domain.Product) error {
	result := config.DB.Save(product)
	return result.Error
}

// DeleteProduct elimina un producto
func (s *ProductService) DeleteProduct(id uint) error {
	result := config.DB.Delete(&domain.Product{}, id)
	return result.Error
}
