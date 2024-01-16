package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toshko07/outdoorsy-challenge/internal/models"
	"github.com/toshko07/outdoorsy-challenge/internal/repositories"
	"go.uber.org/mock/gomock"
)

func TestRetails_GetRetail(t *testing.T) {
	testCases := []struct {
		name                 string
		id                   int
		expectedRepoResponse *models.Rental
		expectedRepoError    error
		expectedRental       *models.Rental
		expectedError        error
	}{
		{
			name:                 "Get existing rental",
			id:                   1,
			expectedRepoResponse: &models.Rental{},
			expectedRepoError:    nil,
			expectedRental:       &models.Rental{},
			expectedError:        nil,
		},
		{
			name:                 "Get non-existing rental",
			id:                   404,
			expectedRepoResponse: nil,
			expectedRepoError:    models.NewNotFoundError("rental with id 404 not found"),
			expectedRental:       nil,
			expectedError:        models.NewNotFoundError("rental with id 404 not found"),
		},
		{
			name:                 "Internal error",
			id:                   500,
			expectedRepoResponse: nil,
			expectedRepoError:    models.NewInternalError("internal error"),
			expectedRental:       nil,
			expectedError:        models.NewInternalError("internal error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			repo := repositories.NewMockRentals(ctrl)
			repo.EXPECT().GetRental(ctx, tc.id).Return(tc.expectedRepoResponse, tc.expectedRepoError)
			service := NewRentalsService(repo)

			// When
			rental, err := service.GetRental(ctx, tc.id)

			// Then
			assert.Equal(t, tc.expectedRental, rental)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestRetails_GetRentals(t *testing.T) {
	testCases := []struct {
		name                 string
		params               models.GetRentalsParams
		expectedRepoResponse []models.Rental
		expectedRepoError    error
		expectedRentals      []models.Rental
		expectedError        error
	}{
		{
			name:                 "Get existing rentals",
			params:               models.GetRentalsParams{},
			expectedRepoResponse: []models.Rental{},
			expectedRepoError:    nil,
			expectedRentals:      []models.Rental{},
			expectedError:        nil,
		},
		{
			name:                 "Get non-existing rentals",
			params:               models.GetRentalsParams{Ids: []int{404}},
			expectedRepoResponse: nil,
			expectedRepoError:    nil,
			expectedRentals:      nil,
			expectedError:        nil,
		},
		{
			name:                 "Internal error",
			params:               models.GetRentalsParams{},
			expectedRepoResponse: nil,
			expectedRepoError:    models.NewInternalError("internal error"),
			expectedRentals:      nil,
			expectedError:        models.NewInternalError("internal error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			repo := repositories.NewMockRentals(ctrl)
			repo.EXPECT().GetRentals(ctx, tc.params).Return(tc.expectedRepoResponse, tc.expectedRepoError)
			service := NewRentalsService(repo)

			// When
			rentals, err := service.GetRentals(ctx, tc.params)

			// Then
			assert.Equal(t, tc.expectedRentals, rentals)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
