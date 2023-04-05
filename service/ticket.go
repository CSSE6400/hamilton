package service

type Ticket struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Concert struct {
		UUID  string `json:"uuid"`
		Name  string `json:"name"`
		Date  string `json:"date"`
		Venue string `json:"venue"`
	}
}
