package main

import "strings"

type CityData struct {
	Name        string
	Connections []Connection
}

type Connection struct {
	cardinalDirection   string
	cityDestinationName string
}

func NewCityData(line string) *CityData {
	lineParts := strings.Fields(line)
	cityConnections := lineParts[1:]
	cityName := lineParts[0]
	cityData := new(CityData)
	cityData.Connections = make([]Connection, len(cityConnections))
	cityData.Name = cityName

	for _, con := range cityConnections {
		cardinalDirection, cityDestinationName := parse(con)
		connect := Connection{cardinalDirection: cardinalDirection, cityDestinationName: cityDestinationName}
		cityData.Connections = append(cityData.Connections, connect)
	}

	return cityData
}

func parse(connection string) (string, string) {
	directionAndDestination := strings.Split(connection, "=")
	return directionAndDestination[0], directionAndDestination[1]
}
