package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/toshko07/outdoorsy-challenge/internal/configs"
	"github.com/toshko07/outdoorsy-challenge/internal/controllers"
	"github.com/toshko07/outdoorsy-challenge/internal/db"
	"github.com/toshko07/outdoorsy-challenge/internal/repositories"
	"github.com/toshko07/outdoorsy-challenge/internal/services"
)

func main() {
	// Setup
	cfg := configs.LoadConfig()
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	// Database
	database := db.Connect(cfg.DB)

	// Repositories
	rentalsRepo := repositories.NewRentalsRepo(database)

	// Services
	rentalsService := services.NewRentalsService(rentalsRepo)

	// Controllers
	rentalsController := controllers.NewRentalsController(rentalsService)

	v1 := e.Group("/v1")
	v1.GET("/rentals/:rental_id", rentalsController.GetRental)
	v1.GET("/rentals", rentalsController.GetRentals)

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", cfg.ServerPort)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatalf("failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
