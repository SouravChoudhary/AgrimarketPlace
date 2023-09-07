package repository

import (
	"agrimarketplace/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ShopRepository defines the interface for interacting with shop data.
type ShopRepository interface {
	FindShopByID(id string) (*models.Shop, error)
	InsertShop(shop *models.Shop) error
	UpdateShop(shop *models.Shop) error
	DeleteShop(id string) error
	FindNearbyShops(latitude, longitude float64, radiusInMeters float64) ([]models.Shop, error)
}

// shopRepository is an implementation of the ShopRepository interface.
type shopRepository struct {
	collection *mongo.Collection
}

// NewShopRepository creates a new instance of the shopRepository.
func NewShopRepository(database *mongo.Database) ShopRepository {
	return &shopRepository{
		collection: database.Collection("shops"),
	}
}

// FindShopByID retrieves a shop by its ID.
func (r *shopRepository) FindShopByID(id string) (*models.Shop, error) {
	var shop models.Shop
	filter := bson.M{"_id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, filter).Decode(&shop)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // Shop not found
		}
		return nil, err // Other error
	}

	return &shop, nil
}

// InsertShop inserts a new shop into the database.
func (r *shopRepository) InsertShop(shop *models.Shop) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, shop)
	if err != nil {
		return err
	}

	return nil
}

// UpdateShop updates an existing shop in the database.
func (r *shopRepository) UpdateShop(shop *models.Shop) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": shop.ID}
	update := bson.M{"$set": shop}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

// DeleteShop deletes a shop by its ID.
func (r *shopRepository) DeleteShop(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

// FindNearbyShops finds nearby shops based on latitude and longitude within a specified radius.
func (r *shopRepository) FindNearbyShops(latitude, longitude float64, radiusInMeters float64) ([]models.Shop, error) {
	// Create a GeoJSON point representing the coordinates
	point := bson.M{
		"type":        "Point",
		"coordinates": []float64{longitude, latitude},
	}

	// Create a GeoJSON query for finding shops within the specified radius
	query := bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry":    point,
				"$maxDistance": radiusInMeters,
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var nearbyShops []models.Shop
	if err := cursor.All(ctx, &nearbyShops); err != nil {
		return nil, err
	}

	return nearbyShops, nil
}
