package services

import (
	"context"

	"github.com/toshko07/outdoorsy-challenge/internal/models"
	"github.com/toshko07/outdoorsy-challenge/internal/repositories"
)

type Rentals interface {
	// GetRental Get Rental by id
	GetRental(ctx context.Context, id int) (*models.Rental, error)
}

type RentalsImpl struct {
	rentalsRepo repositories.Rentals
}

func NewRentalsImpl(rentalsRepo repositories.Rentals) *RentalsImpl {
	return &RentalsImpl{rentalsRepo}
}

// GetRental Get Rental by id
func (r *RentalsImpl) GetRental(ctx context.Context, id int) (*models.Rental, error) {
	return r.rentalsRepo.GetRental(ctx, id)
}
