package repositories

import (
	"MentalHealthCare/database"
	"MentalHealthCare/models"
	"MentalHealthCare/utils"
)


type UserRepositoryImpl struct{}

// func InitUserRepository() UserRepository {
// 	return &UserRepositoryImpl{}
// }

func (ur *UserRepositoryImpl) GetAll() ([]models.User, error) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepositoryImpl) GetByID(id string) (models.User, error) {
	var user models.User

	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) GetByEmail(email string) (models.User, error) {
	var user models.User

	if err := database.DB.First(&user, "email = ?", email).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) Create(userReq models.RegisterRequest) (models.User, error) {
	hashedPassword, err := utils.HashPassword(userReq.Password)
    if err != nil {
        return models.User{}, err
    } // hash the user password 

	var user models.User = models.User{
		Name: userReq.Name,
		Email: userReq.Email,
		Password: hashedPassword,
		Role: userReq.Role,
	}  
	result := database.DB.Create(&user)

	if err := result.Error; err != nil {
		return models.User{}, err
	}

	if err := result.Last(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}


func (ur *UserRepositoryImpl) Update(userReq models.UserUpdateRequest, id string) (models.User, error) {
	user, err := ur.GetByID(id)

	if err != nil {
		return models.User{}, err
	}

	user.Name = userReq.Name
	user.Email = userReq.Email
	user.Password = userReq.Password
	user.Role = userReq.Role

	if err := database.DB.Save(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) Delete(id string) error {
	user, err := ur.GetByID(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
