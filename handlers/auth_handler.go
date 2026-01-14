package handlers

import (
	"myapp/config"
	"myapp/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	FName       string `json:"f_name"`
	LName       string `json:"l_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country_code"`
	Password    string `json:"password"`
}

// Login Request Struct
type LoginRequest struct {
	Phone       string `json:"phone"`
	CountryCode string `json:"country_code"`
	Password    string `json:"password"`
}

func Register(c echo.Context) error {
	req := new(RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request",
		})
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := models.User{
		FName:       req.FName,
		LName:       req.LName,
		Email:       req.Email,
		Phone:       req.Phone,
		CountryCode: req.CountryCode,
		Password:    string(hashedPassword),
	}
	if result := config.DB.Create(&user); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not create user"})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  true,
		"message": "User registered successfully",
		"user":    user,
	})
}

func Login(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Input"})
	}

	var user models.User
	result := config.DB.Where("phone = ? AND country_code = ?", req.Phone, req.CountryCode).First(&user)

	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"status": false, "message": "User not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"status": false, "message": "Invalid password"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 2400).Unix(),
	})

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Token Error"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,

		"token": t,
		"user":  user,
	})
}

func GetProfile(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	var user models.User

	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}

	days := int(time.Since(user.CreatedAt).Hours() / 24)
	user.MemberSinceDays = days
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   user,
	})
}
