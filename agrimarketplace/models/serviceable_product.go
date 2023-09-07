package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ServiceableProduct represents a serviceable product.
type ServiceableProduct struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	ProductID     primitive.ObjectID `bson:"product_id"`
	ShopID        primitive.ObjectID `bson:"shop_id"`
	IsServiceable bool               `bson:"is_serviceable"`
}
