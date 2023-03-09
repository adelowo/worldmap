package alien

import (
	"github.com/Pallinder/go-randomdata"
)

type Alien struct {
	Name             string
	City             string
	CurrentIteration int
}

func NewAlien(city *City) *Alien {
	return &Alien{
		Name:             randomdata.FirstName(randomdata.RandomGender),
		City:             city.Name,
		CurrentIteration: 0,
	}
}
