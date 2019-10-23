package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var alienCount = flag.Int("aliens", 3, "number of aliens to randomly place in cities")

	flag.Parse()

	worldMap := NewWorldMap()
	lines := FileReader("cities.txt")
	worldMap.LoadCities(lines)
	worldMap.LoadAliens(*alienCount)

	simulate(worldMap)

	logRemaindingCities(worldMap)
}

const rounds = 10

func simulate(worldMap *WorldMap) {
	for {
		if canEndSimulation(worldMap) {
			break
		}

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
		_, ok := CacheCityNamesFromInput[name]
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

func canEndSimulation(worldMap *WorldMap) bool {
	return allAliensInactive(worldMap.Aliens) ||
		allAliensReachMaxMoves(worldMap.Aliens) ||
		allAliensTrapped(worldMap.Aliens)
}

func allAliensInactive(aliens []*Alien) bool {
	return allAliens(aliens, func(a *Alien) bool {
		return !a.isActive()
	})
}

func allAliensReachMaxMoves(aliens []*Alien) bool {
	aliens = filterAliens(aliens, func(a *Alien) bool {
		return a.isActive() && !a.isTrapped()
	})

	return allAliens(aliens, func(a *Alien) bool {
		return a.MoveCount >= 10000
	})

}

func allAliensTrapped(aliens []*Alien) bool {
	aliens = filterAliens(aliens, func(a *Alien) bool {
		return a.isActive()
	})

	return allAliens(aliens, func(a *Alien) bool {
		return a.isTrapped()
	})
}

func allAliens(vs []*Alien, f func(*Alien) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func filterAliens(vs []*Alien, f func(*Alien) bool) []*Alien {
	vsf := make([]*Alien, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
