package citydata

import (
	"fmt"
	"strings"
)

var CacheCityNamesFromInput = make(map[string]bool)

type CityData struct {
	Name        string
	Connections []Connection
}

type Connection struct {
	CardinalDirection   string
	CityDestinationName string
}

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

func parse(connection string) ([]string, error) {
	directionAndDestination := strings.Split(connection, "=")
	if len(directionAndDestination) != 2 {
		return nil, fmt.Errorf("invalif format for connection")
	}
	cache := []string{strings.TrimSpace(directionAndDestination[0]), strings.TrimSpace(directionAndDestination[1])}
	return cache, nil
}
