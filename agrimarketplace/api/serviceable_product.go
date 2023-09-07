package api

import (
	"agrimarketplace/service"
	"encoding/json"
	"net/http"
)

// ServiceableProductHandler handles HTTP requests related to serviceable products.
type ServiceableProductHandler struct {
	service service.ServiceableProductService
}

// NewServiceableProductHandler creates a new instance of the ServiceableProductHandler.
func NewServiceableProductHandler(service service.ServiceableProductService) *ServiceableProductHandler {
	return &ServiceableProductHandler{
		service: service,
	}
}

// FindServiceableProductsHandler handles GET requests to retrieve serviceable products.
func (h *ServiceableProductHandler) FindServiceableProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Call the service to find serviceable products
	serviceableProducts, err := h.service.FindServiceableProducts()
	if err != nil {
		http.Error(w, "Failed to retrieve serviceable products", http.StatusInternalServerError)
		return
	}

	// Serialize the serviceable products to JSON
	responseJSON, err := json.Marshal(serviceableProducts)
	if err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseJSON)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
