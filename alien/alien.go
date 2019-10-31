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

func AllAliensInactive(aliens []*Alien) bool {
	return allAliens(aliens, func(a *Alien) bool {
		return !a.IsActive()
	})
}

func AllAliensReachMaxMoves(aliens []*Alien) bool {
	aliens = filterAliens(aliens, func(a *Alien) bool {
		return a.IsActive() && !a.IsTrapped()
	})

	return allAliens(aliens, func(a *Alien) bool {
		return a.MoveCount >= 10000
	})

}

func AllAliensTrapped(aliens []*Alien) bool {
	aliens = filterAliens(aliens, func(a *Alien) bool {
		return a.IsActive()
	})

	return allAliens(aliens, func(a *Alien) bool {
		return a.IsTrapped()
	})
}

func allAliens(vs []*Alien, f func(*Alien) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func filterAliens(vs []*Alien, f func(*Alien) bool) []*Alien {
	vsf := make([]*Alien, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
