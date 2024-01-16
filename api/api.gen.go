// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Error The default error returned
type Error struct {
	// Details The details about the error.
	Details string `json:"details"`

	// Status The HTTP status code returned.
	Status int `json:"status"`

	// Title Error title.
	Title string `json:"title"`
}

// Location The rental location.
type Location struct {
	// City The rental city.
	City string `json:"city"`

	// Country The rental country.
	Country string `json:"country"`

	// Lat The rental latitude.
	Lat float64 `json:"lat"`

	// Lng The rental longitude.
	Lng float64 `json:"lng"`

	// State The rental state.
	State string `json:"state"`

	// Zip The rental zip.
	Zip string `json:"zip"`
}

// Price The rental price.
type Price struct {
	// Day The rental price per day.
	Day int64 `json:"day"`
}

// Rental A rental object.
type Rental struct {
	// Description The rental description.
	Description string `json:"description"`

	// Id The rental id.
	Id int `json:"id"`

	// Length The rental length.
	Length float32 `json:"length"`

	// Location The rental location.
	Location Location `json:"location"`

	// Make The rental make.
	Make string `json:"make"`

	// Model The rental model.
	Model string `json:"model"`

	// Name The rental name.
	Name string `json:"name"`

	// Price The rental price.
	Price Price `json:"price"`

	// PrimaryImageUrl The rental primary image url.
	PrimaryImageUrl string `json:"primary_image_url"`

	// Sleeps The rental sleeps.
	Sleeps int `json:"sleeps"`

	// Type The rental type.
	Type string `json:"type"`

	// User The rental user.
	User User `json:"user"`

	// Year The rental year.
	Year int `json:"year"`
}

// User The rental user.
type User struct {
	// FirstName The rental user first name.
	FirstName string `json:"first_name"`

	// Id The rental
	Id int `json:"id"`

	// LastName The rental user last name.
	LastName string `json:"last_name"`
}

// GetV1RentalsParams defines parameters for GetV1Rentals.
type GetV1RentalsParams struct {
	// PriceMin The minimum price of the rental.
	PriceMin *float32 `form:"price_min,omitempty" json:"price_min,omitempty"`

	// PriceMax The maximum price of the rental.
	PriceMax *float32 `form:"price_max,omitempty" json:"price_max,omitempty"`

	// Limit The maximum number of rentals to return.
	Limit *float32 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset The offset of the rentals to return.
	Offset *float32 `form:"offset,omitempty" json:"offset,omitempty"`

	// Ids The comma separated list of rental ids to return.
	Ids *string `form:"ids,omitempty" json:"ids,omitempty"`

	// Near The comma separated pair [lat,lng] to return rentals near.
	Near *string `form:"near,omitempty" json:"near,omitempty"`

	// Sort The sort order of the rentals to return.
	Sort *string `form:"sort,omitempty" json:"sort,omitempty"`
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xYQW/jNhP9KwS/76jIUuwkiG9pu9umaNHAm+wlCAJaGtvciqSWHG3iDfzfC5KSbNm0",
	"rLQo0FMM83HmcebxceI3milRKgkSDZ2+UZOtQDD38YPWStsPOZhM8xK5knRK71dAcliwqkACFkI0YKUl",
	"5DSipVYlaORg/EZkvDDHYrhFwuaqQoIr8NFiGlF4ZaIswAG5IRq+VmCQvDBD5symwXVpVw1qLpd0E1GD",
	"DKsjiX65v78jHkAylUPLt5NqkiRtXC4RlqBtYORoV/fjutoQt9gl/APLyczzPeS5iag9C9eQ0+ljQ7pJ",
	"ErX1emp3qvkXyNAS+U1lzCcPnVGDRFaQogbFB63IOK57t1pA9yifmCQfNZMZN5kKFT1TlUR9IqzHdCM/",
	"fLoJxSsY9p+OIccq71Z8fBVfXU2uI7pQWtgANFfV3JWzTiArMffNLOTyRPnk8jBDen4eT9LryaAUtqfQ",
	"m8QhuvX4MViO77zsDfSdl50w15M0GZ/SnNNBQ9Pn2DbSt8DXKaTBO82z/sOVFnGovpytT28jJWiSs65W",
	"Unst27pziZcTenhP9w5p04X4z1y2QyY3DQ8PDfDfhfecY2ch4GPcELEms8+hZvO8NzDvmlUasqoC5BJX",
	"/QJ3kE6o8yS+CF2VHbv5v4YFndL/jbZPxah+J0atLW0iKtif/eqwgG5dPiod9HOhcij6Y1nEXrCz9CIJ",
	"RZNM9BOzgG6s3481qmzuQF9V/EXxaMH0+pkLtoTnShenroFFE4cmld473wqxNNPR6OXlJVYV5kpps44z",
	"JUb1xjO3Mfg+FgCl6fclB+m+icEX0X3TE8gC9vytYMaQoMlVBvSpaj4Yn3kNTPdmtoA9cafXJ92CWwU6",
	"iUSd0PXGWteNJmsa7XVraxtqdiOXnftUHzlkTw91LY6ez+48NKcF1wafT2rcbiYOG5D7r2ol329Lpy2J",
	"DWZmoQFiPyk4OUe5Bu5UYTfvYZ3tbi4X6pDSH82dIjd3t2ShNBFMsiWXy5rpdlbbA9OIfgNtfJg0TuLE",
	"nl6VIFnJ6ZSO4yRObeMYrlzPRt/SURNz+kaXEJh8Zm5MNYSRghskatGwsBWyAnCCus3plP4M+DmdtRxL",
	"ppkABG3o9DFUeMElF5Wo3121cOO3j26Dc4v7WoEbCXz7vJCfBbcy8RfTMt4OH8nO/Nw8I5somJy9/s3k",
	"7DWc/Ori/dk9aqeqBFX9n8ExFgUXHMMMxgOzq8XCAHYPPSCz3xZOfTkwdaaEYMSAFQdCvicqwvMBPHhu",
	"wiToOJpEF6GbOoRKybgmjwXDqJDLpy2NtkCyNvYQJ+ntOEhqHF9OorM0vYqvx0PJGaWRKJ17cbyrTXbr",
	"ESrNO7BP4clamSmVNN7Mz5PE/bumJIJ0psDKsuD+8Rh9MX4i22bgCMKcekDrubd9vCnTmq29Fe57zs4c",
	"bDdcvJNPHw3/g0Ig661E0JIV5IP/BcAikC2te9HG1Z7slzu2OXrzH555vhlgobXK5+t6lD7un/7PbT7E",
	"RzvTuVOEdfitIFqKdPfFQl1BUCWBN/QfC2SILgboYJJM/n0dzMCoSmdApEKyUJXM4/+aBrffvjVNblY3",
	"T5u/AgAA///4h37UThMAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
