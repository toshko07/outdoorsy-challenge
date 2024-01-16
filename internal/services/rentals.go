package services

import (
	"context"

	"github.com/toshko07/outdoorsy-challenge/internal/models"
	"github.com/toshko07/outdoorsy-challenge/internal/repositories"
)

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
type Rentals interface {
	GetRental(ctx context.Context, id int) (*models.Rental, error)
	GetRentals(ctx context.Context, params models.GetRentalsParams) ([]models.Rental, error)
}

type RentalsImpl struct {
	rentalsRepo repositories.Rentals
}

func NewRentalsService(rentalsRepo repositories.Rentals) Rentals {
	return &RentalsImpl{rentalsRepo}
}

func (r *RentalsImpl) GetRental(ctx context.Context, id int) (*models.Rental, error) {
	return r.rentalsRepo.GetRental(ctx, id)
}

func (r *RentalsImpl) GetRentals(ctx context.Context, params models.GetRentalsParams) ([]models.Rental, error) {
	return r.rentalsRepo.GetRentals(ctx, params)
}
