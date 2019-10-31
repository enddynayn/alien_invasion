package alien

import (
	"fmt"

	"github.com/enddynayn/alien_invasion/city"
)

// Alien represents all the attributes that an alien has. It also keeps track
// of the number of times an alien has moved.
type Alien struct {
	City      *city.City
	Name      int
	Active    bool
	MoveCount int
}

// NewAlien it returns a new empty alien struct.
func NewAlien() *Alien {
	return new(Alien)
}

// Move moves and alien to another city by randomly selectin a path.
// It returns an error if an alien attempts to move and has nowhere to go.
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

// Deactivate removes an alien from the list of active aliens.
func (a *Alien) Deactivate() error {
	a.Active = false
	return nil
}

// IsActive returns a boolean that check if an alien is active.
func (a *Alien) IsActive() bool {
	return a.Active
}

// IsTrapped checks if an is trapped.
// It returns a boolean.
func (a *Alien) IsTrapped() bool {
	return len(a.City.Paths) == 0
}

// AllAliensInactive  checks if all the aliens are inactive (destroyed).
// It returns a boolean.
func AllAliensInactive(aliens []*Alien) bool {
	return allAliens(aliens, func(a *Alien) bool {
		return !a.IsActive()
	})
}

// AllAliensReachMaxMoves Checks filters all the aliens that are active and not trapped.
// It returns a boolean if all the aliens have moved 1000 times.
func AllAliensReachMaxMoves(aliens []*Alien) bool {
	aliens = filterAliens(aliens, func(a *Alien) bool {
		return a.IsActive() && !a.IsTrapped()
	})

	return allAliens(aliens, func(a *Alien) bool {
		return a.MoveCount >= 10000
	})

}

// AllAliensTrapped checks if all aliens are trapped.
// It returns a boolean.
func AllAliensTrapped(aliens []*Alien) bool {
	aliens = filterAliens(aliens, func(a *Alien) bool {
		return a.IsActive()
	})

	return allAliens(aliens, func(a *Alien) bool {
		return a.IsTrapped()
	})
}

// allAliens check if a collection of aliens all
// have met a condition given a callback.
// It returns a boolean indicating if all aliens met the criteria.
func allAliens(vs []*Alien, f func(*Alien) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// filterAliens filters a collection of aliens given a callback.
// It returns a collection of aliens.
func filterAliens(vs []*Alien, f func(*Alien) bool) []*Alien {
	vsf := make([]*Alien, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
