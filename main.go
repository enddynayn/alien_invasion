package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	worldMap := NewWorldMap()
	lines := []string{"Foo north=Bar west=Baz south=Qu-ux", "Bar south=Foo west=Bee"}
	worldMap.LoadCities(lines)
	worldMap.LoadAliens(5)

	simulate(worldMap)

	logRemaindingCities(worldMap)
}

const rounds = 10

func simulate(worldMap *WorldMap) {
	for i := 0; i < rounds; i++ {

		numberOfAliens := len(worldMap.Aliens)
		m := 0
		for m < numberOfAliens {
			if !worldMap.Aliens[m].Active {
				m++
				continue
			}

			currentAlien := worldMap.Aliens[m]
			currentAlien.Move()

			var aliensInSameCity []int
			n := 0
			for n < numberOfAliens {
				nextAlien := worldMap.Aliens[n]

				if !nextAlien.isActive() {
					n++
					continue
				}

				if n != m && currentAlien.City.Name == nextAlien.City.Name {
					aliensInSameCity = append(aliensInSameCity, nextAlien.Name)
					nextAlien.Deactivate()
				}
				n++
			}

			if len(aliensInSameCity) > 0 {
				aliensInSameCity = append(aliensInSameCity, currentAlien.Name)
				fightLog(currentAlien.City.Name, aliensInSameCity)
				worldMap.RemoveCity(currentAlien.City.Name)
				currentAlien.Deactivate()
			}

			m++
		}
	}
}

func fightLog(cityName string, alienNames []int) {
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

func logRemaindingCities(worldMap *WorldMap) {
	for name, city := range worldMap.Cities {
		_, ok := CacheCitiesFromInputNames[name]
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
