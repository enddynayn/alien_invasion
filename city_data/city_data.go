package citydata

import (
	"fmt"
	"strings"
)

// CacheCityNamesFromInput caches all the all the cities from the input file.
// This is used to log the results in the same format as the input once the simulation
// is complete.
var CacheCityNamesFromInput = make(map[string]bool)

// CityData stores the name of the city and list of connections other cities.
type CityData struct {
	Name        string
	Connections []Connection
}

// Connection stores the cardinal direction and city destination.
type Connection struct {
	CardinalDirection   string
	CityDestinationName string
}

// NewCityData takes in a line from the city data and returns a city
// CityData struct.
func NewCityData(line string) CityData {
	lineParts := strings.Fields(line)
	cityConnections := lineParts[1:]
	cityName := strings.TrimSpace(lineParts[0])
	var cityData CityData

	cityData.Name = cityName
	cityData.Connections = make([]Connection, len(cityConnections))
	CacheCityNamesFromInput[cityData.Name] = true

	for index := range cityData.Connections {
		destination, err := parse(cityConnections[index])
		if err != nil {
			fmt.Println("invalid city connection", cityConnections[index], err)
		}
		cityData.Connections[index] = Connection{CardinalDirection: destination[0], CityDestinationName: destination[1]}
	}
	return cityData
}

// parse splits the a connection and return an splice with
// the first element being the cardinal direction and the second
// element being the destination.
func parse(connection string) ([]string, error) {
	directionAndDestination := strings.Split(connection, "=")
	if len(directionAndDestination) != 2 {
		return nil, fmt.Errorf("invalif format for connection")
	}
	cache := []string{strings.TrimSpace(directionAndDestination[0]), strings.TrimSpace(directionAndDestination[1])}
	return cache, nil
}
