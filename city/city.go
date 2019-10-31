package city

import (
	"fmt"
	"math/rand"
	"time"
)

// City represent a city with name and the paths that connect to other cities.
type City struct {
	Name  string
	Paths map[string]*City
}

// NewCity return a city struct.
func NewCity(name string) *City {
	return &City{Name: name, Paths: make(map[string]*City)}
}

// AddPath adds a bidirectional connection between two cities.
func (c *City) AddPath(cardinalDirection string, cityDestination *City) error {
	c.Paths[cardinalDirection] = cityDestination

	if oppositeDirection, err := OppositeCardinalDirection(cardinalDirection); err == nil {
		cityDestination.Paths[oppositeDirection] = c
	}

	return nil
}

// RemovePath removes a path from one city to the next.
func (c *City) RemovePath(cardinalDirection string) {
	delete(c.Paths, cardinalDirection)
}

// paths returns all the cardinal directions a city contains.
func (c *City) paths() []string {
	keys := make([]string, len(c.Paths))
	i := 0
	for k := range c.Paths {
		keys[i] = k
		i++
	}

	return keys
}

// RandomCityDestination randomly select a cardinal path.
// It returns a radomly selected city.
func (c *City) RandomCityDestination() *City {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	randomNumber := rand.Intn(c.numberOfPaths())

	paths := c.paths()
	randomPath := paths[randomNumber]

	city, _ := c.Paths[randomPath]
	return city
}

// OppositeCardinalDirection returns the opposite cardinal direction
// of a direction.
func OppositeCardinalDirection(cardinalDirection string) (string, error) {
	oppositeCardinalDirections := map[string]string{
		"north": "south",
		"south": "north",
		"east":  "west",
		"west":  "east",
	}

	val, ok := oppositeCardinalDirections[cardinalDirection]

	if !ok {
		return "", fmt.Errorf("no cardinal direction")
	}
	return val, nil
}

// numberOfPaths returns the number of paths a city has.
func (c *City) numberOfPaths() int {
	return len(c.paths())
}
