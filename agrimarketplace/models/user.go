package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the MongoDB database.
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"` // Hashed password
	Email     string             `bson:"email"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Location  string             `bson:"location"`
	Latitude  float64            `bson:"latitude"`
	Longitude float64            `bson:"longitude"`
}
