package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/toshko07/outdoorsy-challenge/internal/models"
	"github.com/toshko07/outdoorsy-challenge/internal/services"
	"go.uber.org/mock/gomock"
)

func TestRentals_GetRental(t *testing.T) {
	testCases := []struct {
		name                    string
		id                      int
		expectedServiceResponse *models.Rental
		expectedServiceError    error
		expectedResponse        string
		expectedStatusCode      int
	}{
		{
			name: "Get existing rental",
			id:   1,
			expectedServiceResponse: &models.Rental{
				Id:              1,
				Name:            "Test Rental",
				Description:     "Test Description",
				Type:            "Test Type",
				Make:            "Test Make",
				Model:           "Test Model",
				Year:            2020,
				Length:          10,
				Sleeps:          2,
				PrimaryImageUrl: "Test Primary Image URL",
				Price:           models.Price{Day: 1},
				Location: models.Location{
					City:    "Test City",
					State:   "Test State",
					Zip:     "Test Zip",
					Country: "Test Country",
					Lat:     19.99,
					Lng:     -19.99,
				},
				User: models.User{Id: 0, FirstName: "Test First Name", LastName: "Test Last Name"},
			},
			expectedServiceError: nil,
			expectedResponse:     "{\"description\":\"Test Description\",\"id\":1,\"length\":10,\"location\":{\"city\":\"Test City\",\"country\":\"Test Country\",\"lat\":19.99,\"lng\":-19.99,\"state\":\"Test State\",\"zip\":\"Test Zip\"},\"make\":\"Test Make\",\"model\":\"Test Model\",\"name\":\"Test Rental\",\"price\":{\"day\":1},\"primary_image_url\":\"Test Primary Image URL\",\"sleeps\":2,\"type\":\"Test Type\",\"user\":{\"first_name\":\"Test First Name\",\"id\":0,\"last_name\":\"Test Last Name\"},\"year\":2020}\n",
			expectedStatusCode:   http.StatusOK,
		},
		{
			name:                    "Get non-existing rental",
			id:                      404,
			expectedServiceResponse: nil,
			expectedServiceError:    models.NewNotFoundError("rental with id '404' not found"),
			expectedResponse:        "{\"details\":\"rental with id '404' not found\",\"status\":404,\"title\":\"Not Found\"}\n",
			expectedStatusCode:      http.StatusNotFound,
		},
		{
			name:                    "Internal server error",
			id:                      500,
			expectedServiceResponse: nil,
			expectedServiceError:    fmt.Errorf("test error"),
			expectedResponse:        "{\"details\":\"internal server error\",\"status\":500,\"title\":\"Internal Server Error\"}\n",
			expectedStatusCode:      http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			ctrl := gomock.NewController(t)
			service := services.NewMockRentals(ctrl)
			service.EXPECT().GetRental(gomock.Any(), tc.id).Return(tc.expectedServiceResponse, tc.expectedServiceError)
			controller := NewRentalsController(service)
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/rentals/:rental_id", nil)
			ctx := e.NewContext(req, rec)
			ctx.SetParamNames("rental_id")
			ctx.SetParamValues(fmt.Sprintf("%d", tc.id))

			// When
			err := controller.GetRental(ctx)

			// Then
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatusCode, rec.Code)
			assert.Equal(t, tc.expectedResponse, rec.Body.String())
		})
	}

}
