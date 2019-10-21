package main

func main() {
	worldMap := NewWorldMap()
	lines := []string{"foo north=bar south=baz east=qu", "bar south=foo east=baz"}
	worldMap.Load(lines)
}
