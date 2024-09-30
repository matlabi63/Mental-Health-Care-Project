package controllers

import (
	"MentalHealthCare/models"
	"MentalHealthCare/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	service services.AdminService
}

func InitAdminController() AdminController {
	return AdminController{
		service: services.InitAdminService(),
	}
}

func (ac *AdminController) GetAll(c echo.Context) error {
	admins, err := ac.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Failed to find admin",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Admin]{
		Status:  "success",
		Message: "All admin",
		Data:    admins,
	})
}

func (ac *AdminController) GetByID(c echo.Context) error {
	adminID := c.Param("id")

	admin, err := ac.service.GetByID(adminID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "Failed",
			Message: "Admin not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Admin]{
		Status:  "success",
		Message: "Admin found",
		Data:    admin,
	})
}

func (ac *AdminController) Create(c echo.Context) error {
	var adminRequest models.AdminRequest

	if err := c.Bind(&adminRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Invalid request",
		})
	}

	err := adminRequest.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Please complete the required fields",
		})
	}

	admin, err := ac.service.Create(adminRequest)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Failed to create an admin",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Admin]{
		Status:  "Success",
		Message: "Admin created",
		Data:    admin,
	})
}

func (ac *AdminController) Update(c echo.Context) error {
	adminID := c.Param("id")

	var adminRequest models.AdminRequest

	if err := c.Bind(&adminRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Invalid request",
		})
	}

	err := adminRequest.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Please complete the required fields",
		})
	}

	admin, err := ac.service.Update(adminRequest, adminID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Failed to update the admin",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Admin]{
		Status:  "success",
		Message: "Admin updated",
		Data:    admin,
	})
}

func (cc *AdminController) Delete(c echo.Context) error {
	adminID := c.Param("id")

	err := cc.service.Delete(adminID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Failed to delete the admin",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "Admin deleted",
	})
}