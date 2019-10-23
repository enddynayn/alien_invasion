package main

import (
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

func (c *City) AddPath(cardinalDirection string, cityDestination *City) bool {
	c.Paths[cardinalDirection] = cityDestination

	//  bidirectional
	if value, ok := OppositeCardinalDirections[cardinalDirection]; ok {
		cityDestination.Paths[value] = c
	}

	return true
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

func (c *City) numberOfPaths() int {
	return len(c.paths())

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

var OppositeCardinalDirections = map[string]string{
	"north": "south",
	"south": "north",
	"east":  "west",
	"west":  "east",
}
