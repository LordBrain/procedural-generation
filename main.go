package main

import (
	"fmt"

	"github.com/blee1170/procedural-generation/cmd"
)

func main() {
	mapSize := 200
	roomSize := 40
	roomNumber := 20
	//MapSize, RoomSize, Max number of rooms.
	town := buildings.NewTown(mapSize, roomSize, roomNumber)
	buildings.PrintRoom(town)
	fmt.Printf("\nRooms: %v\n", len(town.Room))
	town.CreateImage(mapSize)
}
