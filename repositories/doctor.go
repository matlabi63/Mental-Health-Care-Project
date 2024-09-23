package repositories

import (
	"MentalHealthCare/database"
	"MentalHealthCare/models"
)

type DoctorRepositoryImpl struct{}

func InitDoctorRepository() DoctorRepository {
	return &DoctorRepositoryImpl{}
}

func (dr *DoctorRepositoryImpl) GetAll() ([]models.Doctor, error) {
	var doctors []models.Doctor

	if err := database.DB.Find(&doctors).Error; err != nil {
		return nil, err
	}

	return doctors, nil
}

func (dr *DoctorRepositoryImpl) GetByID(id string) (models.Doctor, error) {
	var doctor models.Doctor

	if err := database.DB.First(&doctor, "id = ?", id).Error; err != nil {
		return models.Doctor{}, err
	}

	return doctor, nil
}

func (dr *DoctorRepositoryImpl) Create(doctorReq models.DoctorRequest) (models.Doctor, error) {
	var doctor models.Doctor = models.Doctor{
		UserID: doctorReq.UserID,
		Name: doctorReq.Name,
		Specialty: doctorReq.Specialty,
	}
	result := database.DB.Create(&doctor)

	if err := result.Error; err != nil {
		return models.Doctor{}, err
	}
	if err := result.Last(&doctor).Error; err != nil {
		return models.Doctor{}, err
	}

	return doctor, nil
}

func (dr *DoctorRepositoryImpl) Update(doctorReq models.DoctorRequest, id string) (models.Doctor, error) {
	doctor, err := dr.GetByID(id)

	if err != nil {
		return models.Doctor{}, err
	}

	doctor.Name = doctorReq.Name
	doctor.Specialty = doctorReq.Specialty

	if err := database.DB.Save(&doctor).Error; err != nil {
		return models.Doctor{}, err
	}

	return doctor, nil
}

func (dr *DoctorRepositoryImpl) Delete(id string) error {
	doctor, err := dr.GetByID(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&doctor).Error; err != nil {
		return err
	}

	return nil
}
