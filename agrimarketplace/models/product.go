package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product represents a product in the MongoDB database.
type Product struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	ProductName   string             `bson:"product_name"`
	Description   string             `bson:"description"`
	CategoryID    primitive.ObjectID `bson:"category_id"`
	Price         float64            `bson:"price"`
	StockQuantity int                `bson:"stock_quantity"`
}
