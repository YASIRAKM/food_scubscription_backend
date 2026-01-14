package main

import (
	"myapp/config"
	"myapp/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize DB
	config.ConnectDB()

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS()) // Enable CORS for mobile/web access

	// Setup Routes
	routes.SetupRoutes(e)

	// Start Server
	e.Logger.Fatal(e.Start(":8080"))
}