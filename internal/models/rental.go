package models

type Rental struct {
	Id              int
	Name            string
	Description     string
	Type            string
	Make            string
	Model           string
	Year            int
	Length          float32
	Sleeps          int
	PrimaryImageUrl string
	Price           Price
	Location        Location
	User            User
}

type Price struct {
	PerDay int64
}

type Location struct {
	City    string
	State   string
	Zip     string
	Country string
	Lat     float64
	Lng     float64
}

type User struct {
	Id        int
	FirstName string
	LastName  string
}

type GetRentalsParams struct {
	PriceMin int64
	PriceMax int64
	Limit    int
	Offset   int
	Ids      []int
	Near     []float64
	Sort     string
}
