package service

import (
	"agrimarketplace/models"
	"agrimarketplace/repository"
)

// ProductService defines the interface for working with products.
type ProductService interface {
	GetProductByID(id string) (*models.Product, error)
	GetProductsByCategory(categoryID string) ([]models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id string) error
}

// productService is an implementation of the ProductService interface.
type productService struct {
	productRepo repository.ProductRepository
}

// NewProductService creates a new instance of the productService.
func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

// GetProductByID retrieves a product by its ID.
func (s *productService) GetProductByID(id string) (*models.Product, error) {
	return s.productRepo.FindProductByID(id)
}

// GetProductsByCategory retrieves products by their category ID.
func (s *productService) GetProductsByCategory(categoryID string) ([]models.Product, error) {
	return s.productRepo.FindProductsByCategory(categoryID)
}

// CreateProduct creates a new product.
func (s *productService) CreateProduct(product *models.Product) error {
	return s.productRepo.InsertProduct(product)
}

// UpdateProduct updates an existing product.
func (s *productService) UpdateProduct(product *models.Product) error {
	return s.productRepo.UpdateProduct(product)
}

// DeleteProduct deletes a product by its ID.
func (s *productService) DeleteProduct(id string) error {
	return s.productRepo.DeleteProduct(id)
}
