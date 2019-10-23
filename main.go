package main

import (
	"fmt"
)

func main() {
	worldMap := NewWorldMap()
	lines := []string{"Foo north=Bar west=Baz south=Qu-ux", "Bar south=Foo west=Bee"}
	worldMap.LoadCities(lines)
	worldMap.LoadAliens(2)

	simulate(worldMap)

	fmt.Println("remaining cities", worldMap.Cities)
	fmt.Println("remaining aliens", worldMap.Aliens)
}

const rounds = 100

func simulate(worldMap *WorldMap) {
	for i := 0; i < rounds; i++ {

		numberOfAliens := len(worldMap.Aliens)
		fmt.Println(numberOfAliens)
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
					fmt.Println("hi")
					aliensInSameCity = append(aliensInSameCity, nextAlien.Name)
					nextAlien.Deactivate()
				}
				n++
			}

			if len(aliensInSameCity) > 0 {
				aliensInSameCity = append(aliensInSameCity, currentAlien.Name)
				fmt.Println("destroyed the following aliens", currentAlien.City.Name, aliensInSameCity)
				worldMap.RemoveCity(currentAlien.City.Name)
				currentAlien.Deactivate()
			}

			m++
		}
	}
}
