package main

type Alien struct {
	City *City
	Name int
}

func NewAlien() *Alien {
	return new(Alien)
}
