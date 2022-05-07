package models

import "github.com/siwodevilheart/bookings/internal/forms"

type TmplData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}

type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}
