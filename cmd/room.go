package buildings

import (
	"fmt"
	"math/rand"
	"time"
)

//RoomStruct for building rooms
type RoomStruct struct {
	X1 int
	X2 int
	Y1 int
	Y2 int
}

//LayoutStruct for holding a bunch of rooms
type LayoutStruct struct {
	Room   []RoomStruct
	Blocks [][]int
	Edges  RoomStruct
}
type Point struct {
	X, Y int
}

func NewTown(size int, buildings int) *LayoutStruct {
	rand.Seed(time.Now().UnixNano())
	block := make([][]int, size)
	for i := 0; i < size; i++ {
		block[i] = make([]int, size)
	}
	layout := &LayoutStruct{
		Room:   []RoomStruct{},
		Blocks: block,
		Edges:  RoomStruct{X1: 1, Y1: 1, X2: size, Y2: size},
	}
	for i := 1; i <= buildings; i++ {
		layout.AddRoom()
	}

	return layout
}

//Intersects Rooms. Make sure they do not over lap.
func (room1 *RoomStruct) Intersects(room2 *RoomStruct) bool {
	var x, y bool
	if (room2.X1 >= room1.X1) && (room2.X1 >= room1.X2) {
		x = true
	}
	if (room2.Y1 >= room1.Y1) && (room2.Y1 >= room1.Y2) {
		y = true
	}
	if (x) && (y) {
		fmt.Println("BUUUUUUUUUUUMP")
		return true
	} else {
		return false
	}
}

//AddRoom returns the x/y cords and the center
func (a *LayoutStruct) AddRoom() {
	//Build a room I think.

	//Values Grid
	x := randomInt(4, 51)
	y := randomInt(4, 51)
	width := randomInt(4, 51)
	height := randomInt(4, 51)
	var x1, x2, y1, y2 int

	//Build the walls
	x1 = x
	x2 = x + width
	y1 = y
	y2 = y + height
	newRoom := RoomStruct{X1: x1, X2: x2, Y1: y1, Y2: y2}

	doIntersect := false
	for _, existingRoom := range a.Room {
		if (&newRoom).Intersects(&existingRoom) {
			doIntersect = true
			break
		}
	}
	if !doIntersect {
		a.Room = append(a.Room, newRoom)

	}

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

func (t *LayoutStruct) PrintTown() {
	for _, row := range t.Blocks {
		for _, col := range row {
			if col == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
func (r *RoomStruct) Contains(p Point) bool {
	return p.X >= r.X1 &&
		p.Y >= r.Y1 &&
		p.X <= r.X1+r.X2 &&
		p.Y <= r.Y1+r.X2
}
func (room *LayoutStruct) SetPoint(p Point, cellType int) {
	if room.Edges.Contains(p) {
		room.Blocks[p.X][p.Y] = cellType
	}
}

func (room *LayoutStruct) setBlock(newBlock *RoomStruct) {
	for x := 0; x < newBlock.X2; x++ {
		for y := 0; y < newBlock.Y2; y++ {
			room.SetPoint(Point{X: x + newBlock.X1, Y: y + newBlock.Y1}, 1)
		}
	}
}

func Contains(arr []int, item int) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}
