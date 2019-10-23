package main

type Alien struct {
	City      *City
	Name      int
	Active    bool
	MoveCount int
}

func NewAlien() *Alien {
	return new(Alien)
}

func (a *Alien) Move() {
	if len(a.City.Paths) == 0 {
		return
	}

	from := a.City
	to := from.RandomCityDestination()

	a.City = to
	a.MoveCount++
}

func (a *Alien) Deactivate() bool {
	a.Active = false
	return false
}

func (a *Alien) isActive() bool {
	return a.Active == true
}

func (a *Alien) isTrapped() bool {
	return len(a.City.Paths) == 0
}
