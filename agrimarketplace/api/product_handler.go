package api

import (
	"agrimarketplace/models"
	"agrimarketplace/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	product, err := h.productService.GetProductByID(productID)
	if err != nil {
		http.Error(w, "Error fetching product", http.StatusInternalServerError)
		return
	}

	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	respondWithJSON(w, product, http.StatusOK)
}

func (h *ProductHandler) GetProductsByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["categoryID"]

	products, err := h.productService.GetProductsByCategory(categoryID)
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, products, http.StatusOK)
}

func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	err = h.productService.CreateProduct(&product)
	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, product, http.StatusCreated)
}

func (h *ProductHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	var updatedProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	updatedProduct.ID, _ = primitive.ObjectIDFromHex(productID)

	err = h.productService.UpdateProduct(&updatedProduct)
	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, updatedProduct, http.StatusOK)
}

func (h *ProductHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	err := h.productService.DeleteProduct(productID)
	if err != nil {
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func respondWithJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding response data", http.StatusInternalServerError)
	}
}
