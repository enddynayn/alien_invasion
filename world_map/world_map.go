package worldmap

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/enddynayn/alien_invasion/alien"
	"github.com/enddynayn/alien_invasion/city"
	citydata "github.com/enddynayn/alien_invasion/city_data"
)

type WorldMap struct {
	Cities map[string]*city.City
	Aliens []*alien.Alien
}

func NewWorldMap() *WorldMap {
	return &WorldMap{Cities: make(map[string]*city.City)}
}

func (world *WorldMap) RemoveCity(name string) error {
	city, _ := world.Cities[name]
	for key, c := range city.Paths {
		oppositeDirection, _ := OppositeCardinalDirection(key)
		if _, ok := c.Paths[oppositeDirection]; ok {
			c.RemovePath(oppositeDirection)
		}
	}

	delete(world.Cities, name)
	return nil
}

func (world *WorldMap) LoadCities(lines []string) {
	for _, line := range lines {
		cityData := citydata.NewCityData(line)

		var currentCity *city.City

		if world.hasCity(cityData.Name) {
			currentCity, _ = world.Cities[cityData.Name]
		} else {
			currentCity = city.NewCity(cityData.Name)
			world.addCity(currentCity)
		}

		for _, connection := range cityData.Connections {
			if !world.hasCity(connection.CityDestinationName) {
				city := city.NewCity(connection.CityDestinationName)
				world.addCity(city)
			}
			currentCity.AddPath(connection.CardinalDirection, world.Cities[connection.CityDestinationName])
		}
	}
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

func (world *WorldMap) LoadAliens(count int) error {
	aliens := make([]*alien.Alien, count)
	for i := 0; i < count; i++ {
		alien := alien.NewAlien()
		alien.City = world.randomCity()
		alien.Name = i
		alien.Active = true
		aliens[i] = alien
	}

	world.Aliens = aliens
	return nil
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

func (world *WorldMap) addCity(city *city.City) error {
	world.Cities[city.Name] = city
	return nil
}

func (world *WorldMap) hasCity(cityName string) bool {
	_, ok := world.Cities[cityName]
	return ok
}

func (world *WorldMap) numberOfCities() int {
	return len(world.CityNames())
}

func (world *WorldMap) randomCity() *city.City {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	randomNumber := rand.Intn(world.numberOfCities())

	cities := world.CityNames()
	randomCityName := cities[randomNumber]

	city, _ := world.Cities[randomCityName]
	return city
}
