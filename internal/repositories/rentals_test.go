package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toshko07/outdoorsy-challenge/internal/models"
)

func TestRetailsImpl_GetRetail(t *testing.T) {
	testCases := []struct {
		name           string
		id             int
		expectedRental *models.Rental
		expected       error
	}{
		{
			name: "Get existing rental",
			id:   1,
			expectedRental: &models.Rental{
				Id:              1,
				Name:            "'Abaco' VW Bay Window: Westfalia Pop-top",
				Description:     "ultrices consectetur torquent posuere phasellus urna faucibus convallis fusce sem felis malesuada luctus diam hendrerit fermentum ante nisl potenti nam laoreet netus est erat mi",
				Type:            "camper-van",
				Make:            "Volkswagen",
				Model:           "Bay Window",
				Year:            1978,
				Length:          15,
				Sleeps:          4,
				PrimaryImageUrl: "https://res.cloudinary.com/outdoorsy/image/upload/v1528586451/p/rentals/4447/images/yd7txtw4hnkjvklg8edg.jpg",
				Price:           models.Price{Day: 0},
				Location: models.Location{
					City:    "",
					State:   "",
					Zip:     "",
					Country: "",
					Lat:     0,
					Lng:     0,
				},
				User: models.User{Id: 0, FirstName: "", LastName: ""},
			},
			expected: nil,
		},
		{
			name:           "Get non-existing rental",
			id:             404,
			expectedRental: nil,
			expected:       models.NewNotFoundError("rental with id 404 not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			setupTestData()
			ctx := context.Background()
			repo := NewRentalsRepo(database)

			// When
			rental, err := repo.GetRental(ctx, tc.id)

			// Then
			assert.Equal(t, tc.expected, err)
			assert.Equal(t, tc.expectedRental, rental)
		})
	}

}
