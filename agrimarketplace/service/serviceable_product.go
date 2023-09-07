package service

import (
	"agrimarketplace/models"
	"agrimarketplace/repository"
)

// ServiceableProductService defines the interface for working with serviceable products.
type ServiceableProductService interface {
	FindServiceableProducts() ([]models.ServiceableProduct, error)
}

// serviceableProductService is an implementation of the ServiceableProductService interface.
type serviceableProductService struct {
	serviceableProductRepo repository.ServiceableProductRepository
}

// NewServiceableProductService creates a new instance of the serviceableProductService.
func NewServiceableProductService(serviceableProductRepo repository.ServiceableProductRepository) ServiceableProductService {
	return &serviceableProductService{
		serviceableProductRepo: serviceableProductRepo,
	}
}

// FindServiceableProducts retrieves serviceable products.
func (s *serviceableProductService) FindServiceableProducts() ([]models.ServiceableProduct, error) {
	// Call the repository to find serviceable products
	serviceableProducts, err := s.serviceableProductRepo.FindServiceableProducts()
	if err != nil {
		return nil, err
	}

	// You can implement additional logic here if needed
	return serviceableProducts, nil
}
