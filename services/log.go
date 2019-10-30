package services

import (
	"fmt"
	"strconv"
	"strings"

	citydata "github.com/enddynayn/alien_invasion/city_data"
	worldmap "github.com/enddynayn/alien_invasion/world_map"
)

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
