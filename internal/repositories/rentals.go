package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/toshko07/outdoorsy-challenge/internal/models"
)

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
type Rentals interface {
	GetRental(ctx context.Context, id int) (*models.Rental, error)
}

type RentalsImpl struct {
	db *sql.DB
}

func NewRentalsRepo(db *sql.DB) Rentals {
	return &RentalsImpl{db}
}

func (r *RentalsImpl) GetRental(ctx context.Context, id int) (*models.Rental, error) {
	query := `
		SELECT 
			r.id,
			user_id,
			users.first_name,
			users.last_name,
			name,
			type,
			description,
			sleeps,
			price_per_day,
			home_city,
			home_state,
			home_zip,
			home_country,
			vehicle_make,
			vehicle_model,
			vehicle_year,
			vehicle_length,
			lat,
			lng,
			primary_image_url
		FROM rentals AS r
		JOIN users ON r.user_id = users.id
		WHERE r.id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var rental models.Rental
	var user models.User
	var price models.Price
	var location models.Location

	err := row.Scan(
		&rental.Id,
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&rental.Name,
		&rental.Type,
		&rental.Description,
		&rental.Sleeps,
		&price.Day,
		&location.City,
		&location.State,
		&location.Zip,
		&location.Country,
		&rental.Make,
		&rental.Model,
		&rental.Year,
		&rental.Length,
		&location.Lat,
		&location.Lng,
		&rental.PrimaryImageUrl,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, models.NewNotFoundError(fmt.Sprintf("rental with id %d not found", id))
		}

		return nil, models.NewInternalError(fmt.Sprintf("failed to get rental: %v", err))
	}

	return &rental, nil
}
