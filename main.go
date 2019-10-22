package main

import "fmt"

func main() {
	worldMap := NewWorldMap()
	lines := []string{"foo north=bar south=baz east=qu", "bar south=foo east=baz"}
	worldMap.LoadCities(lines)
	worldMap.LoadAliens(2)
	fmt.Println(worldMap.Aliens)
	fmt.Println(worldMap.Aliens[0].City)
}
