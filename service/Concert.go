package service

type Concert struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Date  string `json:"date"`
	Venue string `json:"venue"`
	Seats struct {
		Max       int `json:"max"`
		Purchased int `json:"purchased"`
	}
}
