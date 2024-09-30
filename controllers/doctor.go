package controllers

import (
	"MentalHealthCare/models"
	"MentalHealthCare/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DoctorController struct {
	service     services.DoctorService
}
 
func InitDoctorController() DoctorController {
	return DoctorController {
		service:     services.InitDoctorService(),
	}
}


func (dc *DoctorController) GetAll(c echo.Context) error {
	doctors, err := dc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Failed with Doctors",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Doctor]{
		Status:  "Success",
		Message: "Found information",
		Data:    doctors,
	})
}

func (dc *DoctorController) GetByID(c echo.Context) error {
	doctorID := c.Param("id")

	doctor, err := dc.service.GetByID(doctorID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "Failed",
			Message: "Doctor not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Doctor]{
		Status:  "Success",
		Message: "Doctor found",
		Data:    doctor,
	})
}

func (dc *DoctorController) Create(c echo.Context) error {
	var doctorReq models.DoctorRequest


	if err := c.Bind(&doctorReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Invalid request",
		})
	}

	err := doctorReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Please complete the required file",
		})
	}

	doctor, err := dc.service.Create(doctorReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Failed to create the doctor",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Doctor]{
		Status:  "Success",
		Message: "Dcotor created",
		Data:    doctor,
	})
}

func (dc *DoctorController) Update(c echo.Context) error {
	var doctorReq models.DoctorRequest

	
	if err := c.Bind(&doctorReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Invalid request",
		})
	}

	doctorID := c.Param("id")

	err := doctorReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Please complete all the required file",
		})
	}

	doctor, err := dc.service.Update(doctorReq, doctorID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Failed to update the Doctors",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Doctor]{
		Status:  "success",
		Message: "Doctor updated",
		Data:    doctor,
	})
}

func (dc *DoctorController) Delete(c echo.Context) error {
	doctorID := c.Param("id")


	err := dc.service.Delete(doctorID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Failed to delete doctor",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "Success",
		Message: "Doctor farm deleted",
	})
}