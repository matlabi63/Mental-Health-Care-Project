package services

import (
	"MentalHealthCare/models"
	"MentalHealthCare/repositories"
	"MentalHealthCare/utils"
)

type AdminService struct {
	Repository repositories.AdminRepository
	JwtOptions models.JWTOptions
}

// NewAdminService initializes and returns an AdminService instance.
// The repository is passed in, which allows for dependency injection.
func NewAdminService(repo repositories.AdminRepository, jwtOptions models.JWTOptions) *AdminService {
	return &AdminService{
		Repository: repo,
		JwtOptions: jwtOptions,
	}
}

// Register creates a new admin account.
func (as *AdminService) Register(adminReq models.Admin) (models.Admin, error) {
	// Delegates admin creation to the repository.
	return as.Repository.Create(adminReq)
}

// Login handles admin login and returns a JWT token if successful.
func (as *AdminService) Login(loginReq models.LoginRequest) (string, error) {
	// Fetch admin by email through repository
	admin, err := as.Repository.GetByEmail(loginReq)
	if err != nil {
		return "", err
	}

	// Generate JWT for admin
	token, err := utils.GenerateJWT(int(admin.ID), as.JwtOptions)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetAdminInfo retrieves an admin's information by ID.
func (as *AdminService) GetAdminInfo(id string) (models.Admin, error) {
	// Fetch admin by ID through repository
	return as.Repository.GetByID(id)
}
