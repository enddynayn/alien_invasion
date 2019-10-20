//https://medium.com/rungo/structures-in-go-76377cc106a2
package main

import (
	"fmt"
	"strings"
)

type WorldMap struct {
	Cities map[string]*City
}

// City should be exported
type City struct {
	Name  string
	Paths map[string]*City
}

var oppositeCardinalDirections = map[string]string{
	"north": "south",
	"south": "north",
	"east":  "west",
	"west":  "east",
}

func main() {
	var worldMap = &WorldMap{Cities: make(map[string]*City)}

	lines := []string{"foo north=bar south=baz east=qu", "bar south=foo east=baz"}

	for _, line := range lines {
		lineParts := strings.Fields(line)
		cityConnections := lineParts[1:]
		cityName := lineParts[0]

		currentCity := &City{Paths: make(map[string]*City)}

		if _, ok := worldMap.Cities[cityName]; ok {
			currentCity = worldMap.Cities[cityName]
		} else {
			currentCity.Name = cityName
			worldMap.Cities[currentCity.Name] = currentCity
		}

		for _, connection := range cityConnections {
			directionAndDestination := strings.Split(connection, "=")
			direction := strings.TrimSpace(directionAndDestination[0])
			destination := strings.TrimSpace(directionAndDestination[1])

			if _, ok := worldMap.Cities[destination]; !ok {
				city := new(City)
				city.Name = destination
				city.Paths = make(map[string]*City)
				worldMap.Cities[destination] = city
			}

			currentCity.Paths[direction] = worldMap.Cities[destination]

			//  bidirectional
			if value, ok := oppositeCardinalDirections[direction]; ok {
				worldMap.Cities[destination].Paths[value] = currentCity
			}
		}

	}

	fmt.Println(worldMap.Cities)
}
