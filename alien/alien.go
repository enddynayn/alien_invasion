package alien

import (
	"fmt"

	"github.com/enddynayn/alien_invasion/city"
)

type Alien struct {
	City      *city.City
	Name      int
	Active    bool
	MoveCount int
}

func NewAlien() *Alien {
	return new(Alien)
}

func (a *Alien) Move() error {
	if len(a.City.Paths) == 0 {
		return fmt.Errorf("alien cannot move")
	}

	from := a.City
	to := from.RandomCityDestination()

	a.City = to
	a.MoveCount++
	return nil
}

func (a *Alien) Deactivate() error {
	a.Active = false
	return nil
}

func (a *Alien) IsActive() bool {
	return a.Active
}

func (a *Alien) IsTrapped() bool {
	return len(a.City.Paths) == 0
}
