package main

import (
	"fmt"

	"github.com/blee1170/procedural-generation/cmd"
)

func main() {
	//MapSize, RoomSize, Max number of rooms.
	town := buildings.NewTown(200, 40, 50)
	buildings.PrintRoom(town)
	fmt.Printf("\nRooms: %v\n", len(town.Room))
	town.CreateImage()
}
