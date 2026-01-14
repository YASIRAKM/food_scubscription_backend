package models

import "time"

type Food struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Type        string    `json:"type"`      // veg, non-veg, eggtarian
	MealTime    string    `json:"meal_time"` // breakfast, lunch, dinner
	CreatedAt   time.Time `json:"created_at"`
}