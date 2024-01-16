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
	GetRentals(ctx context.Context, params models.GetRentalsParams) ([]models.Rental, error)
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
		&price.PerDay,
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

	rental.User = user
	rental.Price = price
	rental.Location = location

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, models.NewNotFoundError(fmt.Sprintf("rental with id %d not found", id))
		}

		return nil, models.NewInternalError(fmt.Sprintf("failed to get rental: %v", err))
	}

	return &rental, nil
}

func (r *RentalsImpl) GetRentals(ctx context.Context, params models.GetRentalsParams) ([]models.Rental, error) {
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
			price_per_day as price,
			home_city as city,
			home_state as state,
			home_zip as zip,
			home_country as country,
			vehicle_make as make,
			vehicle_model as model,
			vehicle_year as year,
			vehicle_length as length,
			lat,
			lng,
			primary_image_url as image_url
		FROM rentals AS r
		JOIN users ON r.user_id = users.id
		WHERE 1 = 1`

	var args []interface{}

	if len(params.Ids) > 0 {
		query += fmt.Sprintf(" AND r.id IN (%s)", createPlaceholders(len(params.Ids)))
		for _, id := range params.Ids {
			args = append(args, id)
		}
	}

	if params.PriceMin > 0 {
		query += fmt.Sprintf(" AND price_per_day >= $%d", len(args)+1)
		args = append(args, params.PriceMin)
	}

	if params.PriceMax > 0 {
		query += fmt.Sprintf(" AND price_per_day <= $%d", len(args)+1)
		args = append(args, params.PriceMax)
	}

	if len(params.Near) > 0 {
		// This code block calculates the distance between two geographical coordinates using the Haversine formula.
		// It filters results within a 100 miles radius.
		// The distance calculation is based on the latitude and longitude values of the coordinates.
		query += fmt.Sprintf(" AND (3959 * acos(cos(radians($%d)) * cos(radians(lat)) * cos(radians(lng) - radians($%d)) + sin(radians($%d)) * sin(radians(lat)))) <= 100", len(args)+1, len(args)+2, len(args)+1)
		args = append(args, params.Near[0], params.Near[1])
	}

	if params.Sort != "" {
		query += fmt.Sprintf(" ORDER BY %s", params.Sort)
	}

	if params.Offset > 0 {
		query += fmt.Sprintf(" OFFSET %d", params.Offset)
	}

	if params.Limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", params.Limit)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, models.NewInternalError(fmt.Sprintf("failed to get rentals: %v", err))
	}

	var rentals []models.Rental

	defer rows.Close()

	for rows.Next() {
		var rental models.Rental
		var user models.User
		var price models.Price
		var location models.Location

		err := rows.Scan(
			&rental.Id,
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&rental.Name,
			&rental.Type,
			&rental.Description,
			&rental.Sleeps,
			&price.PerDay,
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

		rental.User = user
		rental.Price = price
		rental.Location = location

		if err != nil {
			return nil, models.NewInternalError(fmt.Sprintf("failed to get rental: %v", err))
		}

		rentals = append(rentals, rental)
	}

	return rentals, nil
}

func createPlaceholders(count int) string {
	placeholders := ""
	for i := 1; i <= count; i++ {
		placeholders += fmt.Sprintf("$%d,", i)
	}

	return placeholders[:len(placeholders)-1]
}
