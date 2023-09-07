package main

import (
	"agrimarketplace/api"
	"agrimarketplace/repository"
	"agrimarketplace/service"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Initialize MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Update with your MongoDB URI
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	err = client.Connect(nil)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(nil)

	// Create a new MongoDB database instance
	database := client.Database("your_database_name") // Replace with your database name

	// Initialize repositories
	userRepository := repository.NewUserRepository(database)
	shopRepository := repository.NewShopRepository(database)
	productRepository := repository.NewProductRepository(database)
	serviceableProductRepository := repository.NewServiceableProductRepository(database)

	// Initialize services
	userService := service.NewUserService(userRepository)
	shopService := service.NewShopService(shopRepository)
	productService := service.NewProductService(productRepository)
	serviceableProductService := service.NewServiceableProductService(serviceableProductRepository)

	// Initialize handlers
	userHandler := api.NewUserHandler(userService)
	shopHandler := api.NewShopHandler(shopService)
	productHandler := api.NewProductHandler(productService)
	serviceableProductHandler := api.NewServiceableProductHandler(serviceableProductService)

	// Create a router and define routes
	router := http.NewServeMux()

	// Define routes for user-related endpoints
	router.HandleFunc("/users/{username}", userHandler.FindUserByUsername)
	router.HandleFunc("/users/{id}", userHandler.FindUserByID)
	router.HandleFunc("/users/add", userHandler.InsertUser)
	router.HandleFunc("/users/update/{id}", userHandler.UpdateUserHandler)
	// Add routes for other user-related endpoints as needed

	// Define routes for shop-related endpoints
	router.HandleFunc("/shops/{id}", shopHandler.FindShopByIDHandler)
	// Add routes for other shop-related endpoints as needed

	// Define routes for product-related endpoints
	router.HandleFunc("/products/{id}", productHandler.GetProductByIDHandler)
	// Add routes for other product-related endpoints as needed

	// Define routes for serviceable product-related endpoints
	router.HandleFunc("/serviceable-products", serviceableProductHandler.FindServiceableProductsHandler)
	// Add routes for other serviceable product-related endpoints as needed

	// Start the HTTP server
	log.Println("Server started on :8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
