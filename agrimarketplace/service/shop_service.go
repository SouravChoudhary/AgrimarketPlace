package service

import (
	"agrimarketplace/models"
	"agrimarketplace/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ShopService defines the interface for working with shops.
type ShopService interface {
	CreateShop(shop *models.Shop) (*models.Shop, error)
	FindShopByID(id string) (*models.Shop, error)
	UpdateShop(shop *models.Shop) error
	DeleteShop(id string) error
	FindNearbyShops(latitude, longitude float64, radiusInMeters float64) ([]models.Shop, error)
}

// shopService is an implementation of the ShopService interface.
type shopService struct {
	shopRepo repository.ShopRepository
}

// NewShopService creates a new instance of the shopService.
func NewShopService(shopRepo repository.ShopRepository) ShopService {
	return &shopService{
		shopRepo: shopRepo,
	}
}

// CreateShop creates a new shop.
func (s *shopService) CreateShop(shop *models.Shop) (*models.Shop, error) {
	// Implement the logic to create a shop, e.g., validate input, generate ID, etc.
	// You can also add additional business logic here.

	// Ensure that the shop doesn't already exist (you may use a unique constraint on ShopName or other criteria)
	existingShop, err := s.shopRepo.FindShopByID(shop.ID.String())
	if err != nil {
		return nil, err
	}
	if existingShop != nil {
		return nil, ErrShopAlreadyExists
	}

	// Generate an ObjectID for the shop
	shop.ID = primitive.NewObjectID()

	// Call the repository to insert the shop into the database
	if err := s.shopRepo.InsertShop(shop); err != nil {
		return nil, err
	}

	return shop, nil
}

// FindShopByID retrieves a shop by its ID.
func (s *shopService) FindShopByID(id string) (*models.Shop, error) {
	shop, err := s.shopRepo.FindShopByID(id)
	if err != nil {
		return nil, err
	}
	return shop, nil
}

// UpdateShop updates an existing shop.
func (s *shopService) UpdateShop(shop *models.Shop) error {
	// Implement the logic to update a shop, e.g., validate input, handle errors, etc.
	// You can also add additional business logic here.

	// Ensure that the shop to be updated exists
	existingShop, err := s.shopRepo.FindShopByID(shop.ID.Hex())
	if err != nil {
		return err
	}
	if existingShop == nil {
		return ErrShopNotFound
	}

	// Call the repository to update the shop in the database
	if err := s.shopRepo.UpdateShop(shop); err != nil {
		return err
	}

	return nil
}

// DeleteShop deletes a shop by its ID.
func (s *shopService) DeleteShop(id string) error {
	// Implement the logic to delete a shop, e.g., handle errors, etc.
	// You can also add additional business logic here.

	// Ensure that the shop to be deleted exists
	existingShop, err := s.shopRepo.FindShopByID(id)
	if err != nil {
		return err
	}
	if existingShop == nil {
		return ErrShopNotFound
	}

	// Call the repository to delete the shop from the database
	if err := s.shopRepo.DeleteShop(id); err != nil {
		return err
	}

	return nil
}

// FindNearbyShops finds nearby shops based on latitude and longitude within a specified radius.
func (s *shopService) FindNearbyShops(latitude, longitude float64, radiusInMeters float64) ([]models.Shop, error) {
	// Implement the logic to find nearby shops based on coordinates and radius.
	// You can use MongoDB's geospatial queries or another method to find nearby shops.

	nearbyShops, err := s.shopRepo.FindNearbyShops(latitude, longitude, radiusInMeters)
	if err != nil {
		return nil, err
	}

	return nearbyShops, nil
}
