package simulation

import (
	"github.com/enddynayn/alien_invasion/alien"
	"github.com/enddynayn/alien_invasion/services"
	worldmap "github.com/enddynayn/alien_invasion/world_map"
)

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

func canEndSimulation(worldMap *worldmap.WorldMap) bool {
	return allAliensInactive(worldMap.Aliens) ||
		allAliensReachMaxMoves(worldMap.Aliens) ||
		allAliensTrapped(worldMap.Aliens)
}

func allAliensInactive(aliens []*alien.Alien) bool {
	return allAliens(aliens, func(a *alien.Alien) bool {
		return !a.IsActive()
	})
}

func allAliensReachMaxMoves(aliens []*alien.Alien) bool {
	aliens = filterAliens(aliens, func(a *alien.Alien) bool {
		return a.IsActive() && !a.IsTrapped()
	})

	return allAliens(aliens, func(a *alien.Alien) bool {
		return a.MoveCount >= 10000
	})

}

func allAliensTrapped(aliens []*alien.Alien) bool {
	aliens = filterAliens(aliens, func(a *alien.Alien) bool {
		return a.IsActive()
	})

	return allAliens(aliens, func(a *alien.Alien) bool {
		return a.IsTrapped()
	})
}

func allAliens(vs []*alien.Alien, f func(*alien.Alien) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func filterAliens(vs []*alien.Alien, f func(*alien.Alien) bool) []*alien.Alien {
	vsf := make([]*alien.Alien, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
