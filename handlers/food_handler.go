package handlers

import (
	"myapp/config"
	"myapp/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type FoodRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Type        string `json:"type"`      // veg, non-veg, eggtarian
	MealTime    string `json:"meal_time"` // breakfast, lunch, dinner
}

func AddFood(c echo.Context) error {
	req := new(FoodRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Data"})
	}

	food := models.Food{
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
		Type:        req.Type,
		MealTime:    req.MealTime,
		CreatedAt:   time.Now(),
	}

	if err := config.DB.Create(&food).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error adding food"})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  true,
		"message": "Food added successfully",
		"data":    food,
	})
}

func GetFoods(c echo.Context) error {
	var foods []models.Food
	config.DB.Find(&foods)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   foods,
	})
}