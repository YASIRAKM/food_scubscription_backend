package models

import (
	"time"
)

type User struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	FName           string    `json:"f_name"`
	LName           string    `json:"l_name"`
	Email           string    `gorm:"unique" json:"email"`
	Phone           string    `json:"phone"`
	CountryCode     string    `json:"country_code"` // Added for login logic
	Password        string    `json:"-"`            // Hide password in JSON
	Image           string    `json:"image"`
	WalletBalance   float64   `json:"wallet_balance"`
	LoyaltyPoint    int       `json:"loyalty_point"`
	RefCode         string    `json:"ref_code"`
	OrderCount      int       `json:"order_count"`
	MemberSinceDays int       `gorm:"-" json:"member_since_days"` // Calculated field
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"-"`
}