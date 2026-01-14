package handlers

import (
	"myapp/config"
	"myapp/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// GetPrice Logic (Identical to your PHP)
func GetPrices(c echo.Context) error {
	prices := map[string]map[string]int{
		"breakfast": {"veg": 50, "eggtarian": 60, "non-veg": 80},
		"lunch":     {"veg": 80, "eggtarian": 100, "non-veg": 120},
		"dinner":    {"veg": 90, "eggtarian": 110, "non-veg": 140},
	}

	dayPrices := map[string]map[string]int{
		"breakfast": {"veg": 100, "eggtarian": 60, "non-veg": 85},
		"lunch":     {"veg": 80, "eggtarian": 100, "non-veg": 120},
		"dinner":    {"veg": 90, "eggtarian": 110, "non-veg": 140},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":        true,
		"single_prices": prices,
		"21day_prices":  dayPrices,
	})
}

type SubRequest struct {
	StartDate       string   `json:"start_date"`
	EndDate         string   `json:"end_date"`
	TotalDays       int      `json:"total_days"`
	MealTimes       []string `json:"meal_times"`
	DietPreference  string   `json:"diet_preference"`
	ExcludeWeekends bool     `json:"exclude_weekends"`
	TotalPrice      float64  `json:"total_price"`
}

func StoreSubscription(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	req := new(SubRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Data"})
	}
	start, _ := time.Parse("2006-01-02", req.StartDate)
	end, _ := time.Parse("2006-01-02", req.EndDate)
	sub := models.Subscription{
		UserID:          userID,
		StartDate:       start,
		EndDate:         end,
		NumberOfDays:    req.TotalDays,
		MealTiming:      req.MealTimes,
		DietPreference:  req.DietPreference,
		ExcludeWeekends: req.ExcludeWeekends,
		Total:           req.TotalPrice,
		CreatedAt:       time.Now(),
	}

	if err := config.DB.Create(&sub).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error creating subscription"})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  true,
		"message": "Meal plan created successfully",
		"data":    sub,
	})
	
}

func GetSubscriptions(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	var subs []models.Subscription
config.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&subs)
if len(subs) == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  true,
			"message": "Not Subscribed",
			"data":    []string{},
		})
	}
return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Subscribed",
		"data":    subs,
	})
}
	
	