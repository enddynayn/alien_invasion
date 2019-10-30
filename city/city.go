package city

import (
	"fmt"
	"math/rand"
	"time"
)

type City struct {
	Name  string
	Paths map[string]*City
}

func NewCity(name string) *City {
	return &City{Name: name, Paths: make(map[string]*City)}
}

func (c *City) AddPath(cardinalDirection string, cityDestination *City) error {
	c.Paths[cardinalDirection] = cityDestination

	if oppositeDirection, err := OppositeCardinalDirection(cardinalDirection); err == nil {
		cityDestination.Paths[oppositeDirection] = c
	}

	return nil
}

func (c *City) RemovePath(cardinalDirection string) {
	delete(c.Paths, cardinalDirection)
}

func (c *City) paths() []string {
	keys := make([]string, len(c.Paths))
	i := 0
	for k := range c.Paths {
		keys[i] = k
		i++
	}

	return keys
}

func (c *City) RandomCityDestination() *City {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	randomNumber := rand.Intn(c.numberOfPaths())

	paths := c.paths()
	randomPath := paths[randomNumber]

	city, _ := c.Paths[randomPath]
	return city
}

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

func (c *City) numberOfPaths() int {
	return len(c.paths())
}
