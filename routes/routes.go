package routes

import (
	"myapp/handlers"
	"myapp/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	api := e.Group("/api")

	// Public Routes
	api.POST("/register", handlers.Register)
	api.POST("/login", handlers.Login)
	api.GET("/prices", handlers.GetPrices)
	api.GET("/foods", handlers.GetFoods) // Allow viewing foods publicly?

	// Protected Routes (Require JWT)
	protected := api.Group("")
	protected.Use(middleware.IsAuthenticated)

	protected.GET("/profile", handlers.GetProfile)
	protected.POST("/subscription", handlers.StoreSubscription)
	protected.GET("/subscriptions", handlers.GetSubscriptions)
	
	// Admin or Protected route for adding food
	protected.POST("/add-food", handlers.AddFood)
}