package conponent_tests

import (
	"net/http"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/steinfletcher/apitest"
	"github.com/toshko07/outdoorsy-challenge/internal/controllers"
	"github.com/toshko07/outdoorsy-challenge/internal/db"
	"github.com/toshko07/outdoorsy-challenge/internal/repositories"
	"github.com/toshko07/outdoorsy-challenge/internal/services"
)

var echoInstance *echo.Echo

func TestMain(m *testing.M) {
	// Setup Echo
	echoInstance = echo.New()
	echoInstance.HideBanner = true
	echoInstance.Debug = false

	// Database
	testDb, shutdown := db.SetupTestDb()
	db.SetupTestData(testDb)

	// Repositories
	rentalsRepo := repositories.NewRentalsRepo(testDb)

	// Services
	rentalsService := services.NewRentalsService(rentalsRepo)

	// Controllers
	rentalsController := controllers.NewRentalsController(rentalsService)

	v1 := echoInstance.Group("/v1")
	v1.GET("/rentals/:rental_id", rentalsController.GetRental)
	v1.GET("/rentals", rentalsController.GetRentals)

	exitCode := m.Run()
	shutdown()
	os.Exit(exitCode)
}

func TestGetRentalById(t *testing.T) {
	apitest.New().
		Handler(echoInstance).
		Get("/v1/rentals/1").
		Expect(t).
		Body(`{
			"description": "ultrices consectetur torquent posuere phasellus urna faucibus convallis fusce sem felis malesuada luctus diam hendrerit fermentum ante nisl potenti nam laoreet netus est erat mi",
			"id": 1,
			"length": 15,
			"location": {
				"city": "Costa Mesa",
				"country": "US",
				"lat": 33.64,
				"lng": -117.93,
				"state": "CA",
				"zip": "92627"
			},
			"make": "Volkswagen",
			"model": "Bay Window",
			"name": "'Abaco' VW Bay Window: Westfalia Pop-top",
			"price": {
				"day": 16900
			},
			"primary_image_url": "https://res.cloudinary.com/outdoorsy/image/upload/v1528586451/p/rentals/4447/images/yd7txtw4hnkjvklg8edg.jpg",
			"sleeps": 4,
			"type": "camper-van",
			"user": {
				"first_name": "John",
				"id": 1,
				"last_name": "Smith"
			},
			"year": 1978
		}`).
		Status(http.StatusOK).
		End()
}

func TestGetRentalByIdNotFound(t *testing.T) {
	apitest.New().
		Handler(echoInstance).
		Get("/v1/rentals/404").
		Expect(t).
		Body(`{
			"details": "rental with id 404 not found",
			"status": 404,
			"title": "Not Found"
		}`).
		Status(http.StatusNotFound).
		End()
}

func TestGetRentals_WithQueryParams(t *testing.T) {
	apitest.New().
		Handler(echoInstance).
		Get("/v1/rentals").
		Query("near", "33.64,-117.93").
		Query("price_min", "9000").
		Query("price_max", "75000").
		Query("limit", "3").
		Query("offset", "3").
		Query("sort", "price").
		Expect(t).
		Body(`[
			{
				"description": "ultrices consectetur torquent posuere phasellus urna faucibus convallis fusce sem felis malesuada luctus diam hendrerit fermentum ante nisl potenti nam laoreet netus est erat mi",
				"id": 1,
				"length": 15,
				"location": {
					"city": "Costa Mesa",
					"country": "US",
					"lat": 33.64,
					"lng": -117.93,
					"state": "CA",
					"zip": "92627"
				},
				"make": "Volkswagen",
				"model": "Bay Window",
				"name": "'Abaco' VW Bay Window: Westfalia Pop-top",
				"price": {
					"day": 16900
				},
				"primary_image_url": "https://res.cloudinary.com/outdoorsy/image/upload/v1528586451/p/rentals/4447/images/yd7txtw4hnkjvklg8edg.jpg",
				"sleeps": 4,
				"type": "camper-van",
				"user": {
					"first_name": "John",
					"id": 1,
					"last_name": "Smith"
				},
				"year": 1978
			},
			{
				"description": "urna iaculis sed ut porttitor mollis ante cubilia ad felis duis varius mollis nascetur metus faucibus ligula ultricies in faucibus morbi imperdiet auctor morbi torquent",
				"id": 3,
				"length": 16,
				"location": {
					"city": "San Diego",
					"country": "US",
					"lat": 32.83,
					"lng": -117.28,
					"state": "CA",
					"zip": "92037"
				},
				"make": "Volkswagen",
				"model": "Westfalia",
				"name": "1984 Volkswagen Westfalia",
				"price": {
					"day": 18000
				},
				"primary_image_url": "https://res.cloudinary.com/outdoorsy/image/upload/v1504395813/p/rentals/21399/images/nxtwdubpapgpmuc65pd1.jpg",
				"sleeps": 4,
				"type": "camper-van",
				"user": {
					"first_name": "Barry",
					"id": 3,
					"last_name": "Martin"
				},
				"year": 1984
			}
		]`).
		Status(http.StatusOK).
		End()
}
