package repository

import (
	"agrimarketplace/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProductRepository defines the interface for interacting with product data.
type ProductRepository interface {
	FindProductByID(id string) (*models.Product, error)
	FindProductsByCategory(categoryID string) ([]models.Product, error)
	InsertProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id string) error
}

// productRepository is an implementation of the ProductRepository interface.
type productRepository struct {
	collection *mongo.Collection
}

// NewProductRepository creates a new instance of the productRepository.
func NewProductRepository(database *mongo.Database) ProductRepository {
	return &productRepository{
		collection: database.Collection("products"),
	}
}

// FindProductByID retrieves a product by its ID.
func (r *productRepository) FindProductByID(id string) (*models.Product, error) {
	var product models.Product
	filter := bson.M{"_id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // Product not found
		}
		return nil, err // Other error
	}

	return &product, nil
}

// FindProductsByCategory retrieves products by their category ID.
func (r *productRepository) FindProductsByCategory(categoryID string) ([]models.Product, error) {
	var products []models.Product
	filter := bson.M{"category_id": categoryID}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

// InsertProduct inserts a new product into the database.
func (r *productRepository) InsertProduct(product *models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}

	return nil
}

// UpdateProduct updates an existing product in the database.
func (r *productRepository) UpdateProduct(product *models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": product.ID}
	update := bson.M{"$set": product}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

// DeleteProduct deletes a product by its ID.
func (r *productRepository) DeleteProduct(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
