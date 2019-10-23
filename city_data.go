package main

import (
	"strings"
)

var CacheCityNamesFromInput map[string]bool = make(map[string]bool)

type CityData struct {
	Name        string
	Connections []Connection
}

type Connection struct {
	cardinalDirection   string
	cityDestinationName string
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
		cardinalDirection, cityDestinationName := parse(cityConnections[index])
		cityData.Connections[index] = Connection{cardinalDirection: cardinalDirection, cityDestinationName: cityDestinationName}
	}
	return cityData
}

func parse(connection string) (string, string) {
	directionAndDestination := strings.Split(connection, "=")
	return strings.TrimSpace(directionAndDestination[0]), strings.TrimSpace(directionAndDestination[1])
}
