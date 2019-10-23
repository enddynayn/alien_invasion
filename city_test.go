package main

import (
	"testing"
)

func TestRemoveCity(t *testing.T) {
	worldMap := NewWorldMap()
	lines := []string{"Foo north=Bar west=Baz south=Qu-ux", "Bar south=Foo west=Bee"}
	worldMap.LoadCities(lines)

	postApocolypse := NewWorldMap()
	postApocolypseData := []string{"Foo west=Baz south=Qu-ux"}
	postApocolypse.LoadCities(postApocolypseData)

	worldMap.RemoveCity("Bar")

	if _, ok := worldMap.Cities["Bar"]; ok {
		t.Errorf("Expected to remove city Bar but it remains")
	}

	if len(worldMap.Cities["Bee"].Paths) != 0 {
		t.Errorf("Path from Bee to Bar has not been removed")
	}

	if len(worldMap.Cities["Foo"].Paths) != 2 {
		t.Errorf("Expected Foo to have 2 paths but has %q", len(worldMap.Cities["Foo"].Paths))
	}
}
