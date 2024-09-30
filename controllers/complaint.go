package controllers

import (
	"MentalHealthCare/models"
	"net/http"

	"MentalHealthCare/services"

	"github.com/labstack/echo/v4"
)


type ComplaintController struct {
	service     services.ComplaintService
	userService services.UserService
}

func InitComplaintController() ComplaintController {
	return ComplaintController {
		service:     services.InitComplaintService(),
		userService: services.InitUserService(models.JWTOptions{}),
	}
}


func (cc *ComplaintController) GetAll(c echo.Context) error {
	complaints, err := cc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Error with Complaint",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Complaint]{
		Status:  "Success",
		Message: "All Complaint",
		Data:    complaints,
	})
}

func (cc *ComplaintController) GetByID(c echo.Context) error {
	complaintID := c.Param("id")

	complaint, err := cc.service.GetByID(complaintID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "Failed",
			Message: "Complaint not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Complaint]{
		Status:  "Success",
		Message: "Complaint found",
		Data:    complaint,
	})
}

func (cc *ComplaintController) Create(c echo.Context) error {
	var complaintReq models.ComplaintRequest

	if err := c.Bind(&complaintReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Invalid request",
		})
	}

	err := complaintReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Invalid request",
		})
	}

	complaint, err := cc.service.Create(complaintReq)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Please complete the required file",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Complaint]{
		Status:  "Success",
		Message: "Complaint updated",
		Data:    complaint,
	})
}

func (cc *ComplaintController) Update(c echo.Context) error {
	var complaintReq models.ComplaintRequest

	if err := c.Bind(&complaintReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Invalid request",
		})
	}

	complaintID := c.Param("id")

	err := complaintReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Please complete the required file",
		})
	}

	complaint, err := cc.service.Update(complaintReq, complaintID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Failed to update the Complaint",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Complaint]{
		Status:  "Success",
		Message: "Complaint updated",
		Data:   complaint,
	})
}

func (cc *ComplaintController) Delete(c echo.Context) error {
	complaintID := c.Param("id")

	err := cc.service.Delete(complaintID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Failed to delete the Complaint",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "Success",
		Message: "Complaint deleted",
	})
}