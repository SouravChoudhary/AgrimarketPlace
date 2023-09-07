package repository

import (
	"agrimarketplace/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ServiceableProductRepository defines the interface for interacting with serviceable product data.
type ServiceableProductRepository interface {
	FindServiceableProducts() ([]models.ServiceableProduct, error)
}

// serviceableProductRepository is an implementation of the ServiceableProductRepository interface.
type serviceableProductRepository struct {
	collection *mongo.Collection
}

// NewServiceableProductRepository creates a new instance of the serviceableProductRepository.
func NewServiceableProductRepository(database *mongo.Database) ServiceableProductRepository {
	return &serviceableProductRepository{
		collection: database.Collection("serviceable_products"),
	}
}

// FindServiceableProducts finds serviceable products.
func (r *serviceableProductRepository) FindServiceableProducts() ([]models.ServiceableProduct, error) {
	var serviceableProducts []models.ServiceableProduct

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Implement the logic to query and return serviceable products
	cursor, err := r.collection.Find(ctx, bson.M{"is_serviceable": true})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &serviceableProducts); err != nil {
		return nil, err
	}

	return serviceableProducts, nil
}
