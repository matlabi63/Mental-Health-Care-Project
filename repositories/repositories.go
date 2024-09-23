package repositories

import "MentalHealthCare/models"

type UserRepository interface {
	Create(registerReq models.RegisterRequest) (models.User, error)
	GetByID(id string) (models.User, error)
	GetByEmail(email string) (models.User, error) // Add this line
	Update(user models.UserUpdateRequest, id string) (models.User, error)
	Delete(id string) error
	// ... other methods
}

// AdminRepository defines the methods for managing admins.
type AdminRepository interface {
	GetAll() ([]models.Admin, error)
	GetByID(id string) (models.Admin, error)
	GetByEmail(loginReq models.LoginRequest) (models.Admin, error)
	Create(adminReq models.Admin) (models.Admin, error)
	Update(adminReq models.Admin, id string) (models.Admin, error)
	Delete(id string) error
}

// ComplaintRepository defines the methods for managing complaints.
type ComplaintRepository interface {
	GetAll() ([]models.Complaint, error)
	GetByID(id string) (models.Complaint, error)
	Create(complaintReq models.Complaint) (models.Complaint, error)
	Update(complaintReq models.Complaint, id string) (models.Complaint, error)
	Delete(id string) error
}

// DoctorRepository defines the methods for managing doctors.
type DoctorRepository interface {
	GetAll() ([]models.Doctor, error)
	GetByID(id string) (models.Doctor, error)
	Create(doctorReq models.DoctorRequest) (models.Doctor, error)
	Update(doctorReq models.DoctorRequest, id string) (models.Doctor, error)
	Delete(id string) error
}

// InformationRepository defines the methods for managing information.
type InformationRepository interface {
	GetAll() ([]models.Information, error)
	GetByID(id string) (models.Information, error)
	Create(infoReq models.InformationRequest) (models.Information, error)
	Update(infoReq models.InformationRequest, id string) (models.Information, error)
	Delete(id string) error
}

// RecommendationRepository defines the methods for managing recommendations.
type RecommendationRepository interface {
	GetAll() ([]models.Recommendation, error)
	GetByID(id string) (models.Recommendation, error)
	Create(recReq models.RecommendationRequest) (models.Recommendation, error)
	Update(recReq models.RecommendationRequest, id string) (models.Recommendation, error)
	Delete(id string) error
}
