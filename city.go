package main

type City struct {
	Name  string
	Paths map[string]*City
}

func NewCity(name string) *City {
	return &City{Name: name, Paths: make(map[string]*City)}
}

func (c *City) AddPath(cardinalDirection string, cityDestination *City) {
	c.Paths[cardinalDirection] = cityDestination

	//  bidirectional
	if value, ok := oppositeCardinalDirections[cardinalDirection]; ok {
		cityDestination.Paths[value] = c
	}
}

var oppositeCardinalDirections = map[string]string{
	"north": "south",
	"south": "north",
	"east":  "west",
	"west":  "east",
}
