package main

import "fmt"

type Alien struct {
	City *City
	Name int
}

func NewAlien() *Alien {
	return new(Alien)
}

type Entities struct {
	aliens []*Alien
}

func main() {
	worldMap := NewWorldMap()
	lines := []string{"foo north=bar south=baz east=qu", "bar south=foo east=baz"}
	worldMap.Load(lines)
	worldMap.LoadAliens(2)
	// fmt.Println(worldMap.randomCity())
	fmt.Println(worldMap.Aliens)
	fmt.Println(worldMap.Aliens[0].City)
}
