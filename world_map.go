package main

import (
	"fmt"
)

type WorldMap struct {
	Cities map[string]*City
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
