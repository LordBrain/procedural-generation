package main

import (
	"fmt"

	"github.com/blee1170/procedural-generation/cmd"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	mapSize    = kingpin.Flag("map", "Map Size.").Default("1000").Int()
	roomSize   = kingpin.Flag("roomSize", "Size of the rooms.").Default("100").Int()
	roomNumber = kingpin.Flag("roomNumber", "Number of rooms.").Default("20").Int()
)

func main() {
	kingpin.Parse()
	//MapSize, RoomSize, Max number of rooms.
	town := buildings.NewTown(*mapSize, *roomSize, *roomNumber)
	buildings.PrintRoom(town)
	fmt.Printf("\nRooms: %v\n", len(town.Room))
	town.CreateImage(*mapSize)
}
