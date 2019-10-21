package main

import (
	"fmt"
	"math/rand"
	"time"
)

type WorldMap struct {
	Cities map[string]*City
	Aliens []*Alien
}

func NewWorldMap() *WorldMap {
	return &WorldMap{Cities: make(map[string]*City)}
}

func (world *WorldMap) hasCity(cityName string) bool {
	_, ok := world.Cities[cityName]
	return ok
}

func (world *WorldMap) addCity(city *City) {
	world.Cities[city.Name] = city
}

func (world *WorldMap) Load(lines []string) {
	for _, line := range lines {
		cityData := NewCityData(line)

		var currentCity *City

		if world.hasCity(cityData.Name) {
			currentCity, _ = world.Cities[cityData.Name]
		} else {
			currentCity = NewCity(cityData.Name)
			world.addCity(currentCity)
		}

		for _, connection := range cityData.Connections {
			if !world.hasCity(connection.cityDestinationName) {
				city := NewCity(connection.cityDestinationName)
				world.addCity(city)
			}

			currentCity.AddPath(connection.cardinalDirection, world.Cities[connection.cityDestinationName])
		}
	}
	fmt.Println(world.Cities)
}

func (world *WorldMap) CityNames() []string {
	keys := make([]string, len(world.Cities))
	i := 0
	for k := range world.Cities {
		keys[i] = k
		i++
	}

	return keys
}

func (world *WorldMap) randomCity() *City {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	randomNumber := rand.Intn(world.numberOfCities())

	cities := world.CityNames()
	randomCityName := cities[randomNumber]

	city, _ := world.Cities[randomCityName]
	return city
}

func (world *WorldMap) numberOfCities() int {
	return len(world.CityNames())
}

func (world *WorldMap) LoadAliens(count int) {
	aliens := make([]*Alien, count)
	for i := 0; i < count; i++ {
		alien := NewAlien()
		alien.City = world.randomCity()
		alien.Name = i
		aliens[i] = alien
	}

	world.Aliens = aliens
}
