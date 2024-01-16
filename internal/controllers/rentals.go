package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/toshko07/outdoorsy-challenge/api"
	"github.com/toshko07/outdoorsy-challenge/internal/models"
	"github.com/toshko07/outdoorsy-challenge/internal/services"
)

type RentalsController struct {
	RentalsService services.Rentals
}

func NewRentalsController(rentalsService services.Rentals) *RentalsController {
	return &RentalsController{
		RentalsService: rentalsService,
	}
}

// Get Rental by id
func (c *RentalsController) GetRental(e echo.Context) error {
	id := e.Param("rental_id")
	rentalId, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("failed to get rental: %v", err)
		return err
	}

	rental, err := c.RentalsService.GetRental(e.Request().Context(), rentalId)
	if err != nil {
		log.Printf("failed to get rental: %v", err)
		return err
	}

	return e.JSON(http.StatusOK, createRentalResponse(*rental))
}

func createRentalResponse(rental models.Rental) api.Rental {
	return api.Rental{
		Id:              rental.Id,
		Name:            rental.Name,
		Description:     rental.Description,
		Type:            rental.Type,
		Make:            rental.Make,
		Model:           rental.Model,
		Year:            rental.Year,
		Length:          rental.Length,
		Sleeps:          rental.Sleeps,
		PrimaryImageUrl: rental.PrimaryImageUrl,
		Price: api.Price{
			Day: rental.Price.Day,
		},
		Location: api.Location{
			City:    rental.Location.City,
			State:   rental.Location.State,
			Zip:     rental.Location.Zip,
			Country: rental.Location.Country,
			Lat:     rental.Location.Lat,
			Lng:     rental.Location.Lng,
		},
		User: api.User{
			Id:        rental.User.Id,
			FirstName: rental.User.FirstName,
			LastName:  rental.User.LastName,
		},
	}
}
