package services

import (
	"MentalHealthCare/models"
	"MentalHealthCare/repositories"
	"MentalHealthCare/utils"
	"errors"
)

type UserService struct {
	Repository repositories.UserRepository
	JwtOptions models.JWTOptions
}

// InitUserService initializes the user service with JWT options.
func InitUserService(jwtOptions models.JWTOptions) UserService {
	return UserService{
		Repository: &repositories.UserRepositoryImpl{},
		JwtOptions: jwtOptions,
	}
}

// Register registers a new user in the system.
func (us *UserService) Register(registerReq models.RegisterRequest) (models.User, error) {
	return us.Repository.Create(registerReq)
}

// Login authenticates the user and generates a JWT token.
func (us *UserService) Login(loginReq models.LoginRequest) (string, error) {
	user, err := us.Repository.GetByEmail(loginReq.Email)

	if err != nil {
		return "", errors.New("user not found")
	}

	// Verify password
	if !utils.CheckPasswordHash(loginReq.Password, user.Password) {
		return "", errors.New("invalid password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(int(user.ID), us.JwtOptions)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserInfo retrieves user information based on user ID.
func (us *UserService) GetUserInfo(id string) (models.User, error) {
	return us.Repository.GetByID(id)
}

// Update updates the user information.
func (us *UserService) Update(userReq models.UserUpdateRequest, id string) (models.User, error) {
	return us.Repository.Update(userReq, id)
}

// Delete deletes a user account by its ID.
func (us *UserService) Delete(id string) error {
	return us.Repository.Delete(id)
}
