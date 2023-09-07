# AgrimarketPlace
Agri marketplace microservice implemented in golang 
This microservice is designed to provide a platform for managing agricultural products and shops in an online marketplace. It includes functionality for users, shops, products, and serviceable products. This README provides an overview of the microservice, its components, and how to set it up.

## Features
- User Management: Create, retrieve, update, and delete user profiles.
- Shop Management: Manage shop details, including location and operating hours.
- Product Catalog: Maintain a platform-level product catalog with detailed product attributes.
- Product Inventory: Allow shops to manage their product inventory.
- Serviceable Products: Determine and display serviceable products for users based on specific criteria.
- Geospatial Queries: Implement geospatial queries to find nearby shops and users.

## Prerequisites
Before running this microservice, ensure you have the following prerequisites installed:
- Go (version >= 1.15)
- MongoDB (running and accessible)

## Getting Started
1. Clone this repository to your local machine.

2. Install the required Go dependencies:

   ```shell
   go mod tidy
   ```

3. Configure your MongoDB connection in the `cmd/server/main.go` file. Update the MongoDB URI and database name as needed:

   ```go
   clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
   database := client.Database("your_database_name")
   ```

4. Build and run the microservice:

   ```shell
   go run cmd/server/main.go
   ```

The microservice should now be running on port 8080.

## API Endpoints
Here are the available API endpoints provided by this microservice (Not all are implemented yet; WIP ):

- `/users`: User-related endpoints (GET, POST, PUT, DELETE)
- `/shops`: Shop-related endpoints (GET, POST, PUT, DELETE)
- `/products`: Product-related endpoints (GET, POST, PUT, DELETE)
- `/serviceable-products`: Serviceable product-related endpoints (GET)



## License
This project is licensed under the [MIT License](LICENSE).
```

Remember to replace placeholders such as `https://github.com/your/repository.git`, `mongodb://localhost:27017`, `your_database_name`, and other relevant details with your actual project information.
