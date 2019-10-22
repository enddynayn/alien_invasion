package main

type Alien struct {
	City   *City
	Name   int
	Active bool
}

func NewAlien() *Alien {
	return new(Alien)
}

func (a *Alien) Move() {
	from := a.City
	to := from.RandomCityDestination()

	a.City = to
}
