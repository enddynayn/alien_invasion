package main

import (
	"flag"

	"github.com/enddynayn/alien_invasion/services"
	"github.com/enddynayn/alien_invasion/simulation"
	worldmap "github.com/enddynayn/alien_invasion/world_map"
)

func main() {
	var alienCount = flag.Int("aliens", 3, "number of aliens to randomly place in cities")

	flag.Parse()

	worldMap := worldmap.NewWorldMap()
	lines := services.FileReader("cities.txt")

	worldMap.LoadCities(lines)
	worldMap.LoadAliens(*alienCount)

	simulation.Run(worldMap)

	services.LogRemainingCities(worldMap)
}
