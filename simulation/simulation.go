package simulation

import (
	"github.com/enddynayn/alien_invasion/alien"
	"github.com/enddynayn/alien_invasion/services"
	worldmap "github.com/enddynayn/alien_invasion/world_map"
)

// Run starts the simulation of aliens moving from city to city.
// The simulation can be terminated when all aliens have been destroyed (inactive),
// all aliens moved 1,000 times, or aliens becomes trapped.
func Run(worldMap *worldmap.WorldMap) {
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

			resolveAliensInSameCity(currentAlien, worldMap, m)

			m++
		}
	}
}

// ResolveAliensInSameCity checks if an alien is in the same city as another alien.
// if an alien is found to be in the same city it deactivates the aliens (destroys)
// and removes the city in which they collided in.
func resolveAliensInSameCity(currentAlien *alien.Alien, worldMap *worldmap.WorldMap, m int) {
	var aliensInSameCity []int
	n := 0
	numberOfAliens := len(worldMap.Aliens)
	for n < numberOfAliens {
		nextAlien := worldMap.Aliens[n]

		if !nextAlien.IsActive() {
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
		services.LogFight(currentAlien.City.Name, aliensInSameCity)
		worldMap.RemoveCity(currentAlien.City.Name)
		currentAlien.Deactivate()
	}

}

// canEndSimulation checks if the simulation can be terminated.
func canEndSimulation(worldMap *worldmap.WorldMap) bool {
	return alien.AllAliensInactive(worldMap.Aliens) ||
		alien.AllAliensReachMaxMoves(worldMap.Aliens) ||
		alien.AllAliensTrapped(worldMap.Aliens)
}
