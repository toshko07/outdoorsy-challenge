package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-playground/form/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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

// use a single instance of Decoder, it caches struct info
var decoder *form.Decoder

// Get Rental by id
func (c *RentalsController) GetRental(e echo.Context) error {
	id := e.Param("rental_id")
	rentalId, err := strconv.Atoi(id)
	if err != nil {
		log.Errorf("failed to get rental: %v", err)
		return handleError(e, models.NewNotFoundError(fmt.Sprintf("rental with id '%s' not found", id)))
	}

	rental, err := c.RentalsService.GetRental(e.Request().Context(), rentalId)
	if err != nil {
		log.Errorf("failed to get rental: %v", err)
		return handleError(e, err)
	}

	return e.JSON(http.StatusOK, createRentalResponse(*rental))
}

// Get Rentals by query
func (c *RentalsController) GetRentals(e echo.Context) error {
	rentalsQueries := consumeQueryParams(e.QueryParams())
	rentals, err := c.RentalsService.GetRentals(e.Request().Context(), rentalsQueries)
	if err != nil {
		log.Errorf("failed to get rentals: %v", err)
		return handleError(e, err)
	}

	rentalsResponse := make([]api.Rental, len(rentals))
	for i, rental := range rentals {
		rentalsResponse[i] = createRentalResponse(rental)
	}

	return e.JSON(http.StatusOK, rentalsResponse)
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
			Day: rental.Price.PerDay,
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

func handleError(e echo.Context, err error) error {
	switch err.(type) {
	case models.NotFoundError:
		return e.JSON(http.StatusNotFound, api.Error{
			Details: err.Error(),
			Status:  http.StatusNotFound,
			Title:   "Not Found",
		})
	default:
		return e.JSON(http.StatusInternalServerError, api.Error{
			Details: "internal server error",
			Status:  http.StatusInternalServerError,
			Title:   "Internal Server Error",
		})
	}
}

func consumeQueryParams(queryParams url.Values) models.GetRentalsParams {
	rentalsQueries := models.GetRentalsParams{}
	if len(queryParams["price_min"]) > 0 {
		rentalsQueries.PriceMin, _ = strconv.ParseInt(queryParams["price_min"][0], 10, 64)
	}

	if len(queryParams["price_max"]) > 0 {
		rentalsQueries.PriceMax, _ = strconv.ParseInt(queryParams["price_max"][0], 10, 64)
	}

	if len(queryParams["limit"]) > 0 {
		rentalsQueries.Limit, _ = strconv.Atoi(queryParams["limit"][0])
	}

	if len(queryParams["offset"]) > 0 {
		rentalsQueries.Offset, _ = strconv.Atoi(queryParams["offset"][0])
	}

	if len(queryParams["ids"]) > 0 {
		ids := strings.Split(queryParams["ids"][0], ",")
		rentalsQueries.Ids = make([]int, len(ids))
		for i, id := range ids {
			rentalsQueries.Ids[i], _ = strconv.Atoi(id)
		}
	}

	if len(queryParams["near"]) > 0 {
		near := strings.Split(queryParams["near"][0], ",")
		rentalsQueries.Near = make([]float64, len(near))
		for i, n := range near {
			rentalsQueries.Near[i], _ = strconv.ParseFloat(n, 64)
		}
	}

	if len(queryParams["sort"]) > 0 {
		rentalsQueries.Sort = queryParams["sort"][0]
	}

	return rentalsQueries
}
