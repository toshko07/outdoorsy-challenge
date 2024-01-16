package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toshko07/outdoorsy-challenge/internal/models"
)

func TestRetails_GetRetail(t *testing.T) {
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
				Price:           models.Price{PerDay: 16900},
				Location: models.Location{
					City:    "Costa Mesa",
					State:   "CA",
					Zip:     "92627",
					Country: "US",
					Lat:     33.64,
					Lng:     -117.93,
				},
				User: models.User{Id: 1, FirstName: "John", LastName: "Smith"},
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

func TestRetails_GetRentals_ByIDs(t *testing.T) {
	testCases := []struct {
		name           string
		ids            []int
		expected       []models.Rental
		expectedError  error
		expectedLength int
	}{
		{
			name: "Get existing rentals",
			ids:  []int{1, 2},
			expected: []models.Rental{
				{
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
					Price:           models.Price{PerDay: 16900},
					Location:        models.Location{City: "Costa Mesa", State: "CA", Zip: "92627", Country: "US", Lat: 33.64, Lng: -117.93},
					User:            models.User{Id: 1, FirstName: "John", LastName: "Smith"},
				},
				{
					Id: 2, Name: "Maupin: Vanagon Camper",
					Description:     "fermentum nullam congue arcu sollicitudin lacus suspendisse nibh semper cursus sapien quis feugiat maecenas nec turpis viverra gravida risus phasellus tortor cras gravida varius scelerisque",
					Type:            "camper-van",
					Make:            "Volkswagen",
					Model:           "Vanagon Camper",
					Year:            1989,
					Length:          15,
					Sleeps:          4,
					PrimaryImageUrl: "https://res.cloudinary.com/outdoorsy/image/upload/v1498568017/p/rentals/11368/images/gmtye6p2eq61v0g7f7e7.jpg",
					Price:           models.Price{PerDay: 15000},
					Location:        models.Location{City: "Portland", State: "OR", Zip: "97202", Country: "US", Lat: 45.51, Lng: -122.68},
					User:            models.User{Id: 2, FirstName: "Jane", LastName: "Doe"},
				},
			},
			expectedError:  nil,
			expectedLength: 2,
		},
		{
			name:           "Get non-existing rentals",
			ids:            []int{404, 405},
			expected:       nil,
			expectedError:  nil,
			expectedLength: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			ctx := context.Background()
			repo := NewRentalsRepo(database)

			// When
			rentals, err := repo.GetRentals(ctx, models.GetRentalsParams{Ids: tc.ids})

			// Then
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedLength, len(rentals))
			assert.Equal(t, tc.expected, rentals)
		})
	}
}

func TestRetails_GetRentals_ByPrice(t *testing.T) {
	testCases := []struct {
		name           string
		priceMin       int64
		priceMax       int64
		expected       []models.Rental
		expectedError  error
		expectedLength int
	}{
		{
			name:     "Get existing rentals by price",
			priceMin: 1000,
			priceMax: 4000,
			expected: []models.Rental{
				{
					Id:              12,
					Name:            "*ESSENTIAL WORKERS - Pearl - The Maui Camping Cruiser",
					Description:     "malesuada neque velit leo pharetra magnis lectus sapien turpis aenean eu blandit per mi accumsan cursus porta conubia per tellus et morbi dictumst et arcu",
					Type:            "camper-van",
					Make:            "Ford",
					Model:           "Other",
					Year:            2010,
					Length:          17,
					Sleeps:          2,
					PrimaryImageUrl: "https://res.cloudinary.com/outdoorsy/image/upload/v1550269521/p/rentals/108507/images/zlruuz6ll72taorfwjs1.jpg",
					Price:           models.Price{PerDay: 3000},
					Location:        models.Location{City: "Kihei", State: "HI", Zip: "96753", Country: "US", Lat: 20.77, Lng: -156.45},
					User:            models.User{Id: 2, FirstName: "Jane", LastName: "Doe"},
				},
			},
			expectedError:  nil,
			expectedLength: 1,
		},
		{
			name:           "Get non-existing rentals by price",
			priceMin:       100000,
			priceMax:       200000,
			expected:       nil,
			expectedError:  nil,
			expectedLength: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			ctx := context.Background()
			repo := NewRentalsRepo(database)

			// When
			rentals, err := repo.GetRentals(ctx, models.GetRentalsParams{PriceMin: tc.priceMin, PriceMax: tc.priceMax})

			// Then
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedLength, len(rentals))
			assert.Equal(t, tc.expected, rentals)
		})
	}
}

func TestRetails_GetRentals_ByLocation(t *testing.T) {
	testCases := []struct {
		name           string
		near           []float64
		expected       []models.Rental
		expectedError  error
		expectedLength int
	}{
		{
			name: "Get existing rentals by location",
			near: []float64{54.71, -2.81},
			expected: []models.Rental{
				{
					Id:              21,
					Name:            "2013 Peugeot Expert SWB",
					Description:     "sem vitae bibendum hendrerit sapien nulla convallis tempus gravida eu libero litora vulputate tempus nulla ac molestie consequat dictum nisl aptent ligula lacus senectus sagittis",
					Type:            "camper-van",
					Make:            "Peugeot",
					Model:           "Expert SWB",
					Year:            2015,
					Length:          4.8,
					Sleeps:          2,
					PrimaryImageUrl: "https://res.cloudinary.com/outdoorsy/image/upload/v1566292990/p/rentals/137450/images/m1axdiiyampit2da6ufu.jpg",
					Price:           models.Price{PerDay: 9000},
					Location:        models.Location{City: "Cumbria", State: "CMA", Zip: "CA11 9TE", Country: "GB", Lat: 54.72, Lng: -2.88},
					User:            models.User{Id: 1, FirstName: "John", LastName: "Smith"},
				},
			},
			expectedError:  nil,
			expectedLength: 1,
		},
		{
			name:           "Get non-existing rentals by location",
			near:           []float64{42.68, 26.32},
			expected:       nil,
			expectedError:  nil,
			expectedLength: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			ctx := context.Background()
			repo := NewRentalsRepo(database)

			// When
			rentals, err := repo.GetRentals(ctx, models.GetRentalsParams{Near: tc.near})

			// Then
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedLength, len(rentals))
			assert.Equal(t, tc.expected, rentals)
		})
	}
}

func TestRetails_GetRentals_Sort_Offset_Limit(t *testing.T) {
	testCases := []struct {
		name           string
		sort           string
		offset         int
		limit          int
		expected       []models.Rental
		expectedError  error
		expectedLength int
	}{
		{
			name:   "Sort by price, offset 1, limit 2",
			sort:   "price",
			offset: 1,
			limit:  2,
			expected: []models.Rental{
				{
					Id:              9,
					Name:            "Maui \"Alani\" camping car SUBARU IMPREZA 4WD  -Cold AC.",
					Description:     "fermentum torquent hac id tortor conubia litora proin sociosqu congue elit ridiculus fames velit viverra faucibus eleifend sagittis etiam aptent sociosqu taciti metus iaculis quam",
					Type:            "camper-van",
					Make:            "SUBARU IMPREZA 4WD",
					Model:           "SUBARU IMPREZA 4WD",
					Year:            2003,
					Length:          13,
					Sleeps:          2,
					PrimaryImageUrl: "https://res.cloudinary.com/outdoorsy/image/upload/v1538027810/p/rentals/82458/images/bphrohl2r4wxc8wg3v11.jpg",
					Price:           models.Price{PerDay: 5900},
					Location:        models.Location{City: "Kahului", State: "HI", Zip: "96732", Country: "US", Lat: 20.88, Lng: -156.45},
					User:            models.User{Id: 4, FirstName: "Todd", LastName: "Edison"},
				},
				{
					Id:              13,
					Name:            "The Coolest Camper Van Around",
					Description:     "porta eros bibendum cum bibendum purus aliquet dis augue litora tempus ridiculus ornare tempor nascetur tristique mauris aenean vehicula maecenas facilisi sociis ut parturient vel",
					Type:            "camper-van",
					Make:            "Dodge",
					Model:           "B Van",
					Year:            2000,
					Length:          16,
					Sleeps:          4,
					PrimaryImageUrl: "https://res.cloudinary.com/outdoorsy/image/upload/v1556142483/p/rentals/109101/images/ea2vvbovq0tvouj00fad.jpg",
					Price:           models.Price{PerDay: 7900},
					Location:        models.Location{City: "Provo", State: "UT", Zip: "84601", Country: "US", Lat: 40.24, Lng: -111.7},
					User:            models.User{Id: 3, FirstName: "Barry", LastName: "Martin"},
				},
			},
			expectedError:  nil,
			expectedLength: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			ctx := context.Background()
			repo := NewRentalsRepo(database)

			// When
			rentals, err := repo.GetRentals(ctx, models.GetRentalsParams{Sort: tc.sort, Offset: tc.offset, Limit: tc.limit})

			// Then
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedLength, len(rentals))
			assert.Equal(t, tc.expected, rentals)
		})
	}
}

func TestRetails_GetRentals_FilterByAllQueries(t *testing.T) {
	testCases := []struct {
		name           string
		ids            []int
		priceMin       int64
		priceMax       int64
		near           []float64
		sort           string
		offset         int
		limit          int
		expected       []models.Rental
		expectedError  error
		expectedLength int
	}{
		{
			name:     "Get existing rentals by ids, near by location, sort by year, offset 2, limit 2",
			ids:      []int{1, 3, 7, 15, 23},
			priceMin: 9000,
			priceMax: 75000,
			near:     []float64{33.64, -117.93},
			sort:     "year",
			offset:   2,
			limit:    3,
			expected: []models.Rental{
				{
					Id: 3, Name: "1984 Volkswagen Westfalia",
					Description:     "urna iaculis sed ut porttitor mollis ante cubilia ad felis duis varius mollis nascetur metus faucibus ligula ultricies in faucibus morbi imperdiet auctor morbi torquent",
					Type:            "camper-van",
					Make:            "Volkswagen",
					Model:           "Westfalia",
					Year:            1984,
					Length:          16,
					Sleeps:          4,
					PrimaryImageUrl: "https://res.cloudinary.com/outdoorsy/image/upload/v1504395813/p/rentals/21399/images/nxtwdubpapgpmuc65pd1.jpg",
					Price:           models.Price{PerDay: 18000},
					Location:        models.Location{City: "San Diego", State: "CA", Zip: "92037", Country: "US", Lat: 32.83, Lng: -117.28},
					User:            models.User{Id: 3, FirstName: "Barry", LastName: "Martin"},
				},
				{
					Id: 7, Name: "2002 Volkswagen Eurovan Weekender Westfalia",
					Description:     "purus neque pellentesque potenti posuere molestie vivamus urna faucibus class justo porta litora turpis cubilia sit class torquent ullamcorper netus ut sapien libero consequat quisque",
					Type:            "camper-van",
					Make:            "VW",
					Model:           "Eurovan Weekender Westfalia",
					Year:            2002,
					Length:          0,
					Sleeps:          4,
					PrimaryImageUrl: "https://res.cloudinary.com/outdoorsy/image/upload/v1526614056/p/rentals/52210/images/nou2lx0h0dsjzbqeotuf.jpg",
					Price:           models.Price{PerDay: 15000},
					Location:        models.Location{City: "Rancho Mission Viejo", State: "CA", Zip: "", Country: "US", Lat: 33.53, Lng: -117.63},
					User:            models.User{Id: 2, FirstName: "Jane", LastName: "Doe"},
				},
				{
					Id: 23, Name: "2002 Chevrolet Van Conversion",
					Description:     "magnis interdum morbi faucibus habitasse sapien porta iaculis platea mi proin posuere vel ligula curabitur amet vehicula amet condimentum ridiculus diam diam proin est etiam",
					Type:            "camper-van",
					Make:            "Chevrolet",
					Model:           "Express",
					Year:            2002,
					Length:          21,
					Sleeps:          2,
					PrimaryImageUrl: "https://res.cloudinary.com/outdoorsy/image/upload/v1569722222/p/rentals/143740/images/ooxoce0zrlycj5esm3jh.png",
					Price:           models.Price{PerDay: 9900},
					Location:        models.Location{City: "San Diego", State: "CA", Zip: "92107", Country: "US", Lat: 32.73, Lng: -117.24},
					User:            models.User{Id: 3, FirstName: "Barry", LastName: "Martin"},
				},
			},
			expectedError:  nil,
			expectedLength: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			ctx := context.Background()
			repo := NewRentalsRepo(database)

			// When
			rentals, err := repo.GetRentals(ctx, models.GetRentalsParams{Ids: tc.ids, PriceMin: tc.priceMin, PriceMax: tc.priceMax, Near: tc.near, Sort: tc.sort, Offset: tc.offset, Limit: tc.limit})

			// Then
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedLength, len(rentals))
			assert.Equal(t, tc.expected, rentals)
		})
	}
}
