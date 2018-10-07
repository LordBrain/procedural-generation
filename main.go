package main

import (
	"math/rand"
	"time"

	"github.com/blee1170/procedural-generation/cmd"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	// mapSize    = kingpin.Flag("map", "Map Size.").Default("1000").Int()
	// roomSize   = kingpin.Flag("roomSize", "Size of the rooms.").Default("100").Int()
	// roomNumber = kingpin.Flag("roomNumber", "Number of rooms.").Default("20").Int()
	// rooms      = kingpin.Flag("rooms", "Do random Rooms").Bool()
	defaultCommand = kingpin.Command("empty", "Default Stuff").Default()
	seed           = kingpin.Command("seed", "Use a seed")
	seedValue      = seed.Arg("seedValue", "Seed Value").Required().Int()
)

func main() {
	switch kingpin.Parse() {
	case "seed":
		buildings.Test(*seedValue)
	case "empty":
		rand.Seed(time.Now().UnixNano())
		newSeed := rand.Intn(100)
		buildings.Test(newSeed)
	}
	//MapSize, RoomSize, Max number of rooms.
	// if *rooms {
	// 	town := buildings.NewTown(*mapSize, *roomSize, *roomNumber)
	// 	buildings.PrintRoom(town)
	// 	fmt.Printf("\nRooms: %v\n", len(town.Room))
	// 	town.CreateImage(*mapSize)
	// }

	// if *seed {
	// 	buildings.Test(seed)
	// } else {
	// 	rand.Seed(time.Now().UnixNano())
	// 	newSeed := rand.Intn(100)
	// 	buildings.Test(newSeed)
	// }

}
