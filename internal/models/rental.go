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
	Day int
}

type Location struct {
	City    string
	State   string
	Zip     string
	Country string
	Lat     float32
	Lng     float32
}

type User struct {
	Id        int
	FirstName string
	LastName  string
}
