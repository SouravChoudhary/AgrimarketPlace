package api

import (
	"agrimarketplace/models"
	"agrimarketplace/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ShopHandler handles HTTP requests related to shops.
type ShopHandler struct {
	ShopService service.ShopService
}

// NewShopHandler creates a new instance of ShopHandler.
func NewShopHandler(shopService service.ShopService) *ShopHandler {
	return &ShopHandler{
		ShopService: shopService,
	}
}

// CreateShopHandler handles the creation of a new shop.
func (h *ShopHandler) CreateShopHandler(w http.ResponseWriter, r *http.Request) {
	var shop models.Shop

	if err := json.NewDecoder(r.Body).Decode(&shop); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdShop, err := h.ShopService.CreateShop(&shop)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdShop)
}

// FindShopByIDHandler handles the retrieval of a shop by ID.
func (h *ShopHandler) FindShopByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shopID := vars["id"]

	shop, err := h.ShopService.FindShopByID(shopID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shop)
}

// UpdateShopHandler handles the update of an existing shop.
func (h *ShopHandler) UpdateShopHandler(w http.ResponseWriter, r *http.Request) {
	var shop models.Shop

	if err := json.NewDecoder(r.Body).Decode(&shop); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.ShopService.UpdateShop(&shop); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteShopHandler handles the deletion of a shop by ID.
func (h *ShopHandler) DeleteShopHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shopID := vars["id"]

	if err := h.ShopService.DeleteShop(shopID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// FindNearbyShopsHandler handles the retrieval of nearby shops.
func (h *ShopHandler) FindNearbyShopsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract latitude, longitude, and radius from request query parameters
	latStr := r.URL.Query().Get("latitude")
	lonStr := r.URL.Query().Get("longitude")
	radiusStr := r.URL.Query().Get("radius")

	// Parse latitude, longitude, and radius as float64
	latitude, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}

	longitude, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		http.Error(w, "Invalid longitude", http.StatusBadRequest)
		return
	}

	radius, err := strconv.ParseFloat(radiusStr, 64)
	if err != nil {
		http.Error(w, "Invalid radius", http.StatusBadRequest)
		return
	}

	// Call the ShopService to find nearby shops
	nearbyShops, err := h.ShopService.FindNearbyShops(latitude, longitude, radius)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nearbyShops)
}
