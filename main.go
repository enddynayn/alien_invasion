package main

import (
	"fmt"
)

func main() {
	worldMap := NewWorldMap()
	lines := []string{"foo north=bar south=baz east=qu", "bar south=foo east=baz"}
	worldMap.LoadCities(lines)
	worldMap.LoadAliens(2)
	fmt.Println(worldMap.Aliens)
	fmt.Println(worldMap.Aliens[0].City)
	simulate(worldMap)

	fmt.Println("destroyed the following aliens", worldMap.Cities)
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

			n := 0
			var aliensInSameCity []int
			var willDeleteCity bool
			for n < numberOfAliens {
				if !worldMap.Aliens[n].Active {
					n++
					continue
				}
				nextAlien := worldMap.Aliens[n]

				if n != m && currentAlien.Name == nextAlien.Name {
					aliensInSameCity = append(aliensInSameCity, nextAlien.Name)
					worldMap.Aliens[n].Active = false
					willDeleteCity = true
				}
				n++

			}

			if willDeleteCity {
				aliensInSameCity = append(aliensInSameCity, currentAlien.Name)
				fmt.Println("destroyed the following aliens", currentAlien.City.Name, aliensInSameCity)
				worldMap.RemoveCity(currentAlien.City)
				currentAlien.Active = false
			}

			m++
		}
	}
}
