package service

import (
	"agrimarketplace/models"
	"agrimarketplace/repository"
)

// UserService defines the interface for working with user data.
type UserService interface {
	FindUserByID(id string) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
	InsertUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
	FindNearbyUsers(latitude, longitude float64, radiusInMeters float64) ([]models.User, error)
}

// userService is an implementation of the UserService interface.
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new instance of the userService.
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// FindUserByID retrieves a user by their ID.
func (s *userService) FindUserByID(id string) (*models.User, error) {
	return s.userRepo.FindUserByID(id)
}

// FindUserByUsername retrieves a user by their username.
func (s *userService) FindUserByUsername(username string) (*models.User, error) {
	return s.userRepo.FindUserByUsername(username)
}

// InsertUser inserts a new user into the database.
func (s *userService) InsertUser(user *models.User) error {
	return s.userRepo.InsertUser(user)
}

// UpdateUser updates an existing user in the database.
func (s *userService) UpdateUser(user *models.User) error {
	return s.userRepo.UpdateUser(user)
}

// DeleteUser deletes a user by their ID.
func (s *userService) DeleteUser(id string) error {
	return s.userRepo.DeleteUser(id)
}

// FindNearbyUsers finds nearby users based on latitude and longitude within a specified radius.
func (s *userService) FindNearbyUsers(latitude, longitude float64, radiusInMeters float64) ([]models.User, error) {
	return s.userRepo.FindNearbyUsers(latitude, longitude, radiusInMeters)
}
