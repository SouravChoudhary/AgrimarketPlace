package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Shop represents a shop in the MongoDB database.
type Shop struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	ShopName       string             `bson:"shop_name"`
	OwnerID        primitive.ObjectID `bson:"owner_id"`
	Location       string             `bson:"location"`
	OperatingHours string             `bson:"operating_hours"`
	Latitude       float64            `bson:"latitude"`
	Longitude      float64            `bson:"longitude"`
}
