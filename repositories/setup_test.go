package repositories_test

import (
	"MentalHealthCare/repositories/mocks"
	"MentalHealthCare/services"
	"os"
	"testing"
)

var (
	userRepository       mocks.UserRepository
	userService          services.UserService

	adminRepository      mocks.AdminRepository
	adminService         services.AdminService

	complaintRepository  mocks.ComplaintRepository
	complaintService     services.ComplaintService

	doctorRepository     mocks.DoctorRepository
	doctorService        services.DoctorService

	informationRepository mocks.InformationRepository
	informationService    services.InformationService

	recommendationRepository mocks.RecommendationRepository
	recommendationService    services.RecommendationService
)

func TestMain(m *testing.M) {
	// userService = services.InitUserService(models.JWTOptions{})
	// userService.Repository = &userRepository

	// adminService = services.InitAdminService()
	// adminService.Repository = &adminRepository

	complaintService = services.InitComplaintService()
	complaintService.Repository = &complaintRepository

	doctorService = services.InitDoctorService()
	doctorService.Repository = &doctorRepository

	informationService = services.InitInformationService()
	informationService.Repository = &informationRepository

	recommendationService = services.InitRecommendationService()
	recommendationService.Repository = &recommendationRepository

	os.Exit(m.Run())
}
