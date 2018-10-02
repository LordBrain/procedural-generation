package buildings

import (
	"fmt"
	"math"
	"math/rand"
)

//RoomStruct for building rooms
type RoomStruct struct {
	X1     int
	X2     int
	Y1     int
	Y2     int
	Center [2]float64
}

//LayoutStruct for holding a bunch of rooms
type LayoutStruct struct {
	Room []RoomStruct
}

//AddRoom returns the x/y cords and the center
func (a *LayoutStruct) AddRoom() {
	//Build a room I think.
	// newRoom := new(RoomStruct)

	//Values Grid
	x := randomInt(4, 51)
	y := randomInt(4, 51)
	width := randomInt(4, 51)
	height := randomInt(4, 51)
	var x1, x2, y1, y2 int
	//Center
	var center [2]float64

	//Build the walls
	x1 = x
	x2 = x + width
	y1 = y
	y2 = y + height
	center[0] = math.Floor((float64(x1) + float64(x2)) / 2)
	center[1] = math.Floor((float64(y1) + float64(y2)) / 2)
	newRoom := RoomStruct{X1: x1, X2: x2, Y1: y1, Y2: y2, Center: center}

	a.Room = append(a.Room, newRoom)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

//PrintRoom displays the rooms made.
func PrintRoom(r *LayoutStruct) {
	numRooms := len(r.Room)

	for i := 0; i < numRooms; i++ {

		topWall := r.Room[i].X2 - r.Room[i].X1
		sideGap := topWall - 2
		sideLenght := r.Room[i].Y2 - r.Room[i].Y1
		topGap := sideLenght - 2

		for i := 0; i <= topWall; i++ {
			fmt.Print("#")
		}
		fmt.Println()
		for x := 0; x <= topGap; x++ {
			fmt.Print("#")
			for y := 0; y <= sideGap; y++ {
				fmt.Print(" ")
			}
			fmt.Print("#")
			fmt.Println()
		}
		for i := 0; i <= topWall; i++ {
			fmt.Print("#")
		}
		fmt.Printf("\n%v\n", r.Room[i])
	}
}
