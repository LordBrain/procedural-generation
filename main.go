package main

import (
	"math/rand"
	"time"

	"github.com/blee1170/procedural-generation/cmd"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	layout := &buildings.LayoutStruct{
		Room: []buildings.RoomStruct{},
	}
	for i := 1; i <= 5; i++ {
		layout.AddRoom()
		// layout.Room = room
		// fmt.Println(layout.Room)
		// fmt.Printf("Room: %v\nX1: %v, X2: %v, Y1: %v, Y2: %v, Center: %v-%v\n", i, room.X1, room.X2, room.Y1, room.Y2, room.Center[0], room.Center[1])

	}
	buildings.PrintRoom(layout)
}
