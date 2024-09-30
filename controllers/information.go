package controllers

import (
	"MentalHealthCare/models"
	"MentalHealthCare/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type InformationController struct {
	service     services.InformationService
}
 
func InitInformationController() InformationController {
	return InformationController {
		service:     services.InitInformationService(),
}
}

func (mc *InformationController) GetAll(c echo.Context) error {
	informations, err := mc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Failed with information",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Information]{
		Status:  "Success",
		Message: "All information",
		Data:    informations,
	})
}

func (mc *InformationController) GetByID(c echo.Context) error {
	informationID := c.Param("id")

	information, err := mc.service.GetByID(informationID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "Failed",
			Message: "Information not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Information]{
		Status:  "Success",
		Message: "Information found",
		Data:    information,
	})
}

func (mc *InformationController) Create(c echo.Context) error {
	var infoReq models.InformationRequest

	if err := c.Bind(&infoReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Invalid request",
		})
	}

	err := infoReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Please complete all the required files",
		})
	}

	information, err := mc.service.Create(infoReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Failed to create the Information",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Information]{
		Status:  "Success",
		Message: "Information created",
		Data:    information,
	})
}

func (mc *InformationController) Update(c echo.Context) error {
	var infoReq models.InformationRequest

	if err := c.Bind(&infoReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Invalid request",
		})
	}

	informationID := c.Param("id")

	err := infoReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Please complete all the required files",
		})
	}

	information, err := mc.service.Update(infoReq, informationID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Failed to update the Information",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Information]{
		Status:  "Success",
		Message: "Information updated",
		Data:    information,
	})
}

func (mc *InformationController) Delete(c echo.Context) error {
	informationID := c.Param("id")


	err := mc.service.Delete(informationID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Failed to delete information",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "Success",
		Message: "Information deleted",
	})
}