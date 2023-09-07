package repository

import (
	"agrimarketplace/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository defines the interface for interacting with user data.
type UserRepository interface {
	FindUserByID(id string) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
	InsertUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
	FindNearbyUsers(latitude, longitude float64, radiusInMeters float64) ([]models.User, error)
}

// userRepository is an implementation of the UserRepository interface.
type userRepository struct {
	collection *mongo.Collection
}

// NewUserRepository creates a new instance of the userRepository.
func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		collection: database.Collection("users"),
	}
}

// FindUserByID retrieves a user by their ID.
func (r *userRepository) FindUserByID(id string) (*models.User, error) {
	var user models.User
	filter := bson.M{"_id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

// FindUserByUsername retrieves a user by their username.
func (r *userRepository) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	filter := bson.M{"username": username}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

// InsertUser inserts a new user into the database.
func (r *userRepository) InsertUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser updates an existing user in the database.
func (r *userRepository) UpdateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes a user by their ID.
func (r *userRepository) DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

// FindNearbyUsers finds nearby users based on latitude and longitude within a specified radius.
func (r *userRepository) FindNearbyUsers(latitude, longitude float64, radiusInMeters float64) ([]models.User, error) {
	// Create a GeoJSON point representing the coordinates
	point := bson.M{
		"type":        "Point",
		"coordinates": []float64{longitude, latitude},
	}

	// Create a GeoJSON query for finding users within the specified radius
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

	var nearbyUsers []models.User
	if err := cursor.All(ctx, &nearbyUsers); err != nil {
		return nil, err
	}

	return nearbyUsers, nil
}
