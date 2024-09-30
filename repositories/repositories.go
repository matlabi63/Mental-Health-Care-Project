package repositories

import (
	"MentalHealthCare/models"
)

//  userRepository define user
type UserRepository interface {
	Register(registerReq models.RegisterRequest) (models.User, error)
	GetByEmail(loginReq models.LoginRequest) (models.User, error)
	GetUserInfo(id string) (models.User, error)
}

// AdminRepository defines admins.
type AdminRepository interface {
	GetAll() ([]models.Admin, error)
	GetByID(id string) (models.Admin, error)
	GetByEmail(loginReq models.LoginRequest) (models.Admin, error)
	Create(adminReq models.AdminRequest) (models.Admin, error)
	Update(adminReq models.AdminRequest, id string) (models.Admin, error)
	Delete(id string) error
}

// ComplaintRepository defines complaints.
type ComplaintRepository interface {
	GetAll() ([]models.Complaint, error)
	GetByID(id string) (models.Complaint, error)
	Create(complaintReq models.ComplaintRequest) (models.Complaint, error)
	Update(complaintReq models.ComplaintRequest, id string) (models.Complaint, error)
	Delete(id string) error
}

// DoctorRepository defines doctors.
type DoctorRepository interface {
	GetAll() ([]models.Doctor, error)
	GetByID(id string) (models.Doctor, error)
	Create(doctorReq models.DoctorRequest) (models.Doctor, error)
	Update(doctorReq models.DoctorRequest, id string) (models.Doctor, error)
	Delete(id string) error
}

// InformationRepository defines information.
type InformationRepository interface {
	GetAll() ([]models.Information, error)
	GetByID(id string) (models.Information, error)
	Create(infoReq models.InformationRequest) (models.Information, error)
	Update(infoReq models.InformationRequest, id string) (models.Information, error)
	Delete(id string) error
}

// RecommendationRepository defines recommendations.
type RecommendationRepository interface {
	GetAll() ([]models.Recommendation, error)
	GetByID(id string) (models.Recommendation, error)
	Create(recReq models.RecommendationRequest) (models.Recommendation, error)
	Update(recReq models.RecommendationRequest, id string) (models.Recommendation, error)
	Delete(id string) error
}
