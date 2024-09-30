package controllers

import (
	"MentalHealthCare/models"
	"MentalHealthCare/services"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RecommendationController struct {
	service     services.RecommendationService
}

func InitRecommendationController() RecommendationController {
	return RecommendationController {
		service:     services.InitRecommendationService(),
	}
}


func (rc *RecommendationController) GetAll(c echo.Context) error {
	recommendations, err := rc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Failed with recommendation",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Recommendation]{
		Status:  "Success",
		Message: "All recommendations",
		Data:    recommendations,
	})
}

func (rc *RecommendationController) GetByID(c echo.Context) error {
	recomendationID := c.Param("id")

	recommendation, err := rc.service.GetByID(recomendationID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "Failed",
			Message: "Recommendation not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Recommendation]{
		Status:  "Success",
		Message: "Recommendation found",
		Data:    recommendation,
	})
}

// func (rc *RecommendationController) Create(c echo.Context) error {
// 	var recomandationReq models.RecommendationRequest

// 	if err := c.Bind(&recomandationReq); err != nil {
// 		return c.JSON(http.StatusBadRequest, models.Response[string]{
// 			Status:  "Failed",
// 			Message: "Invalid request",
// 		})
// 	}

// 	err := recomandationReq.Validate()

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, models.Response[string]{
// 			Status:  "Failed",
// 			Message: "Please complete all the required files",
// 		})
// 	}

// 	recommendation, err := rc.service.Create(recomandationReq)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, models.Response[string]{
// 			Status:  "Failed",
// 			Message: "Failed to create the recommendation",
// 		})
// 	}

// 	return c.JSON(http.StatusCreated, models.Response[models.Recommendation]{
// 		Status:  "Success",
// 		Message: "Recommendation created",
// 		Data:    recommendation,
// 	})
// }

func (rc *RecommendationController) Create(c echo.Context) error {
	var recommendationReq models.RecommendationRequest

	// Bind the request body to the recommendationReq struct
	if err := c.Bind(&recommendationReq); err != nil {
			log.Println("Bind error:", err)
			return c.JSON(http.StatusBadRequest, models.Response[string]{
					Status:  "Failed",
					Message: "Invalid request",
			})
	}

	// Validate the recommendation request
	if err := recommendationReq.Validate(); err != nil {
			log.Println("Validation error:", err)
			return c.JSON(http.StatusBadRequest, models.Response[string]{
					Status:  "Failed",
					Message: "Please complete all the required fields",
			})
	}

	// Call the service layer to create a recommendation
	recommendation, err := rc.service.Create(recommendationReq)
	if err != nil {
			log.Println("Service error:", err)
			return c.JSON(http.StatusInternalServerError, models.Response[string]{
					Status:  "Failed",
					Message: "Failed to create the recommendation",
			})
	}

	// Successful response
	return c.JSON(http.StatusOK, models.Response[models.Recommendation]{
			Status:  "Success",
			Message: "Recommendation created successfully",
			Data:    recommendation,
	})
}


func (rc *RecommendationController) Update(c echo.Context) error {
	var recomandationReq models.RecommendationRequest

	if err := c.Bind(&recomandationReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	recommendationID := c.Param("id")

	err := recomandationReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "Failed",
			Message: "Please complete all the required files",
		})
	}

	recommendation, err := rc.service.Update(recomandationReq, recommendationID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Failed to update the recommendation",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Recommendation]{
		Status:  "Success",
		Message: "Recommendation updated",
		Data:    recommendation,
	})
}

func (rc *RecommendationController) Delete(c echo.Context) error {
	recommendationID := c.Param("id")


	err := rc.service.Delete(recommendationID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "Failed",
			Message: "Failed to delete Recommendation",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "Success",
		Message: "Recommendation deleted",
	})
}