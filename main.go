package main

import (
	"fmt"

	"github.com/blee1170/procedural-generation/cmd"
)

func main() {

	town := buildings.NewTown(50, 20)
	buildings.PrintRoom(town)
	town.PrintTown()
	fmt.Printf("\nRooms: %v\n", len(town.Room))
}
