package services

import (
	"fmt"
	"strconv"
	"strings"

	citydata "github.com/enddynayn/alien_invasion/city_data"
	worldmap "github.com/enddynayn/alien_invasion/world_map"
)

// LogFight logs the fight between aliens.
// It outputs the city it destroy with the names
// of aliens that destroy it. For example,
// Foo has been destroyed by alien 0 and alien 1!
func LogFight(cityName string, alienNames []int) {
	format := make([]string, 0)

	for index, alienName := range alienNames {
		name := strconv.FormatInt(int64(alienName), 10)

		if index == len(alienNames)-1 {
			format = append(format, name)
			continue
		}
		format = append(format, name+" and alien")

	}
	output := cityName + " has been destroyed by alien " + strings.TrimSpace(strings.Join(format, " ")) + "!"

	fmt.Println(output)
}

// LogRemainingCities it logs the remaining cities
// after after a simulation has ended. For example,
// Bar west=Bee
func LogRemainingCities(worldMap *worldmap.WorldMap) {
	for name, city := range worldMap.Cities {
		_, ok := citydata.CacheCityNamesFromInput[name]
		if !ok {
			continue
		}

		var line string

		line = name

		for cardinalDirection, cityDestination := range city.Paths {
			line = line + " " + cardinalDirection + "=" + cityDestination.Name
		}

		fmt.Println(line)
	}
}
