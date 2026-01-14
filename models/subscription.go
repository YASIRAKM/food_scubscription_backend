package models

import (
	"time"

	"github.com/lib/pq" // Driver for String Array support if using raw array, or use serializer
)

type Subscription struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	UserID          uint           `json:"user_id"`
	StartDate       time.Time      `json:"start_date"`
	EndDate         time.Time      `json:"end_date"`
	NumberOfDays    int            `json:"number_of_days"`
	MealTiming      pq.StringArray `gorm:"type:text[]" json:"meal_timing"` // Stores ["breakfast", "lunch"]
	DietPreference  string         `json:"diet_preference"`
	ExcludeWeekends bool           `json:"exclude_weekends"`
	Total           float64        `json:"total"`
	CreatedAt       time.Time      `json:"created_at"`
}