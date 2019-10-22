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

			var aliensInSameCity []int
			n := 0
			for n < numberOfAliens {
				nextAlien := worldMap.Aliens[n]

				if !nextAlien.isActive() {
					n++
					continue
				}

				if n != m && currentAlien.Name == nextAlien.Name {
					aliensInSameCity = append(aliensInSameCity, nextAlien.Name)
					nextAlien.Deactivate()
				}
				n++
			}

			if len(aliensInSameCity) > 0 {
				aliensInSameCity = append(aliensInSameCity, currentAlien.Name)
				fmt.Println("destroyed the following aliens", currentAlien.City.Name, aliensInSameCity)
				worldMap.RemoveCity(currentAlien.City)
				currentAlien.Deactivate()
			}

			m++
		}
	}
}
